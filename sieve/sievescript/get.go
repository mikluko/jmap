package sievescript

import (
	"github.com/mikluko/jmap"
	"github.com/mikluko/jmap/sieve"
)

type Get struct {
	Account             jmap.ID              `json:"accountId,omitempty"`
	IDs                 []jmap.ID            `json:"ids,omitempty"`
	Properties          []string             `json:"properties,omitempty"`
	ReferenceIDs        *jmap.ResultReference `json:"#ids,omitempty"`
	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}

func (m *Get) Name() string         { return "SieveScript/get" }
func (m *Get) Requires() []jmap.URI { return []jmap.URI{sieve.URI} }

type GetResponse struct {
	Account  jmap.ID        `json:"accountId,omitempty"`
	State    string         `json:"state,omitempty"`
	List     []*SieveScript `json:"list,omitempty"`
	NotFound []string       `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
