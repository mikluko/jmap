package sievescript

import (
	"github.com/mikluko/jmap"
	"github.com/mikluko/jmap/sieve"
)

type QueryChanges struct {
	Account         jmap.ID           `json:"accountId,omitempty"`
	Filter          Filter            `json:"filter,omitempty"`
	Sort            []*SortComparator `json:"sort,omitempty"`
	SinceQueryState string            `json:"sinceQueryState,omitempty"`
	MaxChanges      uint64            `json:"maxChanges,omitempty"`
	UpToID          jmap.ID           `json:"upToId,omitempty"`
	CalculateTotal  bool              `json:"calculateTotal,omitempty"`
}

func (m *QueryChanges) Name() string         { return "SieveScript/queryChanges" }
func (m *QueryChanges) Requires() []jmap.URI { return []jmap.URI{sieve.URI} }

type QueryChangesResponse struct {
	Account       jmap.ID           `json:"accountId,omitempty"`
	OldQueryState string            `json:"oldQueryState,omitempty"`
	NewQueryState string            `json:"newQueryState,omitempty"`
	Removed       []jmap.ID         `json:"removed,omitempty"`
	Added         []*jmap.AddedItem `json:"added,omitempty"`
}

func newQueryChangesResponse() jmap.MethodResponse { return &QueryChangesResponse{} }
