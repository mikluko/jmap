package push

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sarak.sh/mikluko/jmap"
)

// stream serves an event stream that writes events, then holds the
// connection open for hold, or closes it at once when hold is zero.
func stream(t *testing.T, hold time.Duration, events ...string) *jmap.Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Email", r.URL.Query().Get("types"))
		assert.Equal(t, "identity", r.Header.Get("Accept-Encoding"))
		w.Header().Set("Content-Type", "text/event-stream")
		for _, e := range events {
			fmt.Fprint(w, e)
		}
		w.(http.Flusher).Flush()
		select {
		case <-r.Context().Done():
		case <-time.After(hold):
		}
	}))
	t.Cleanup(srv.Close)
	return &jmap.Client{
		HttpClient: srv.Client(),
		Session:    &jmap.Session{EventSourceURL: srv.URL + "/es?types={types}&closeafter={closeafter}&ping={ping}"},
	}
}

const stateEvent = "event: state\ndata: {\"@type\":\"StateChange\",\"changed\":{\"a\":{\"Email\":\"s2\"}}}\n\n"

func TestListenContextDeliversStatesThenReportsClose(t *testing.T) {
	var got []string
	var connected bool
	es := &EventSource{
		Client:    stream(t, 0, ": comment\n\n", stateEvent, "event: ping\ndata: {}\n\n"),
		Events:    []jmap.EventType{"Email"},
		OnConnect: func() { connected = true },
		Handler: func(sc *jmap.StateChange) {
			assert.True(t, connected, "OnConnect runs before the first event")
			got = append(got, sc.Changed["a"]["Email"])
		},
	}
	err := es.ListenContext(context.Background())
	require.ErrorIs(t, err, ErrClosed)
	assert.Equal(t, []string{"s2"}, got)
}

func TestListenContextReportsSilence(t *testing.T) {
	es := &EventSource{
		Client:  stream(t, time.Minute),
		Events:  []jmap.EventType{"Email"},
		Ping:    1,
		Handler: func(*jmap.StateChange) {},
	}
	began := time.Now()
	err := es.ListenContext(context.Background())
	require.ErrorIs(t, err, ErrSilent)
	assert.Less(t, time.Since(began), 10*time.Second)
}

func TestListenContextEndsWithItsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	es := &EventSource{
		Client:    stream(t, time.Minute),
		Events:    []jmap.EventType{"Email"},
		OnConnect: cancel,
		Handler:   func(*jmap.StateChange) {},
	}
	require.ErrorIs(t, es.ListenContext(ctx), context.Canceled)
}

func TestListenKeepsReturningNilOnClose(t *testing.T) {
	es := &EventSource{Client: stream(t, 0), Events: []jmap.EventType{"Email"}, Handler: func(*jmap.StateChange) {}}
	assert.NoError(t, es.Listen())
}
