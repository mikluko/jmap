package sievescript

import (
	"github.com/mikluko/jmap"
	"github.com/mikluko/jmap/sieve"
)

type Query struct {
	Account        jmap.ID           `json:"accountId,omitempty"`
	Filter         Filter            `json:"filter,omitempty"`
	Sort           []*SortComparator `json:"sort,omitempty"`
	Position       int64             `json:"position,omitempty"`
	Anchor         jmap.ID           `json:"anchor,omitempty"`
	AnchorOffset   int64             `json:"anchorOffset,omitempty"`
	Limit          uint64            `json:"limit,omitempty"`
	CalculateTotal bool              `json:"calculateTotal,omitempty"`
}

func (m *Query) Name() string         { return "SieveScript/query" }
func (m *Query) Requires() []jmap.URI { return []jmap.URI{sieve.URI} }

type QueryResponse struct {
	Account             jmap.ID   `json:"accountId,omitempty"`
	QueryState          string    `json:"queryState,omitempty"`
	CanCalculateChanges bool      `json:"canCalculateChanges,omitempty"`
	Position            uint64    `json:"position,omitempty"`
	IDs                 []jmap.ID `json:"ids,omitempty"`
	Total               int64     `json:"total,omitempty"`
	Limit               uint64    `json:"limit,omitempty"`
}

func newQueryResponse() jmap.MethodResponse { return &QueryResponse{} }
