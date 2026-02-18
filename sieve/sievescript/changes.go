package sievescript

import (
	"github.com/mikluko/jmap"
	"github.com/mikluko/jmap/sieve"
)

type Changes struct {
	Account    jmap.ID `json:"accountId,omitempty"`
	SinceState string  `json:"sinceState,omitempty"`
	MaxChanges uint64  `json:"maxChanges,omitempty"`
}

func (m *Changes) Name() string         { return "SieveScript/changes" }
func (m *Changes) Requires() []jmap.URI { return []jmap.URI{sieve.URI} }

type ChangesResponse struct {
	Account        jmap.ID   `json:"accountId,omitempty"`
	OldState       string    `json:"oldState,omitempty"`
	NewState       string    `json:"newState,omitempty"`
	HasMoreChanges bool      `json:"hasMoreChanges,omitempty"`
	Created        []jmap.ID `json:"created,omitempty"`
	Updated        []jmap.ID `json:"updated,omitempty"`
	Destroyed      []jmap.ID `json:"destroyed,omitempty"`
}

func newChangesResponse() jmap.MethodResponse { return &ChangesResponse{} }
