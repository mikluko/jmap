package push

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"sarak.sh/mikluko/jmap"
)

var (
	// ErrClosed is ListenContext's report that the server ended the stream.
	ErrClosed = errors.New("push: event stream closed")
	// ErrSilent is ListenContext's report that the stream sent nothing,
	// not even a ping, for twice the Ping interval.
	ErrSilent = errors.New("push: event stream went silent")
)

// A subscription to an event stream
type EventSource struct {
	// The JMAP client to use for the stream
	Client *jmap.Client

	// The function to pass state change events to
	Handler func(*jmap.StateChange)

	// The events to subscribe to. If left unset, will default to AllEvents
	Events []jmap.EventType

	// Interval the server should ping the client at, in seconds. The server
	// may choose to ignore this value. Set to 0 to disable pinging (which
	// the server may also ignore)
	Ping uint

	// Whether to close the connection after a state event
	CloseAfterState bool

	// OnConnect, if set, is called by ListenContext once the server has
	// accepted the stream and before the first event is read. Changes made
	// before the stream opened reach no Handler, so a client that keeps
	// state catches up here.
	OnConnect func()

	// The response of the request.
	resp *http.Response
}

// Connect to the server
func (e *EventSource) connect(ctx context.Context) error {
	if e.Client.Session == nil {
		if err := e.Client.Authenticate(); err != nil {
			return err
		}
	}
	// Create the URL for the subscription
	url, err := url.Parse(e.Client.Session.EventSourceURL)
	if err != nil {
		return err
	}
	q := url.Query()

	if len(e.Events) == 0 {
		e.Events = []jmap.EventType{jmap.AllEvents}
	}
	// types field
	types := []string{}
	for _, e := range e.Events {
		types = append(types, string(e))
	}
	typeStr := strings.Join(types, ",")
	q.Set("types", typeStr)

	// ping field
	q.Set("ping", fmt.Sprintf("%d", e.Ping))

	// close after field
	closeAfter := "no"
	if e.CloseAfterState {
		closeAfter = "state"
	}
	q.Set("closeafter", closeAfter)

	// set the query string
	url.RawQuery = q.Encode()

	// make the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "text/event-stream")
	// A compressed stream reaches the client only in compressor-sized
	// chunks: Stalwart 1.0.3 sends not even the response headers until the
	// first event. Naming the encoding also stops Go's transport from
	// asking for gzip.
	req.Header.Set("Accept-Encoding", "identity")
	e.resp, err = e.Client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if e.resp.StatusCode != 200 {
		e.resp.Body.Close()
		return fmt.Errorf("invalid request, response code: %d", e.resp.StatusCode)
	}
	return nil
}

// Starts listening for events from the source. Listen will block when called
// and return when the source has been disconnected or closed via a call to
// Close()
func (e *EventSource) Listen() error {
	err := e.ListenContext(context.Background())
	if errors.Is(err, ErrClosed) {
		return nil
	}
	return err
}

// ListenContext listens like Listen, until ctx ends or the stream does. It
// returns ctx's error, ErrClosed when the server or Close ends the stream,
// or ErrSilent when Ping is set and nothing arrives for twice its interval.
func (e *EventSource) ListenContext(ctx context.Context) error {
	parent := ctx
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := e.connect(ctx); err != nil {
		return err
	}
	defer e.resp.Body.Close()

	alive := make(chan struct{}, 1)
	if e.Ping > 0 {
		go watchdog(ctx, cancel, alive, 2*time.Duration(e.Ping)*time.Second)
	}
	if e.OnConnect != nil {
		e.OnConnect()
	}

	scanner := bufio.NewScanner(e.resp.Body)
	scanner.Buffer(make([]byte, 64*1024), 4*1024*1024)
	var event string

	for scanner.Scan() {
		select {
		case alive <- struct{}{}:
		default:
		}
		// https://html.spec.whatwg.org/multipage/server-sent-events.html#event-stream-interpretation
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, ":"):
			// ignore the line, it's a comment
		case strings.Contains(line, ":"):
			fields := strings.SplitN(line, ":", 2)
			if len(fields) < 2 {
				continue
			}
			k := fields[0]
			v := strings.TrimSpace(fields[1])
			switch k {
			case "event":
				event = v
			case "data":
				switch event {
				case "state":
					state := &jmap.StateChange{}
					err := json.Unmarshal([]byte(v), state)
					if err != nil {
						return err
					}
					e.Handler(state)
				}
			}
		}
	}
	switch {
	case parent.Err() != nil:
		return parent.Err()
	case ctx.Err() != nil:
		return ErrSilent
	}
	return ErrClosed
}

// watchdog cancels when alive receives nothing for silence.
func watchdog(ctx context.Context, cancel func(), alive <-chan struct{}, silence time.Duration) {
	t := time.NewTimer(silence)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-alive:
			t.Reset(silence)
		case <-t.C:
			cancel()
			return
		}
	}
}

// Closes the stream
func (e *EventSource) Close() {
	if e.resp != nil && e.resp.Body != nil {
		e.resp.Body.Close()
	}
}
