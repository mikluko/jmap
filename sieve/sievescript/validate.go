package sievescript

import (
	"github.com/mikluko/jmap"
	"github.com/mikluko/jmap/sieve"
)

type Validate struct {
	Account jmap.ID `json:"accountId,omitempty"`
	BlobID  jmap.ID `json:"blobId,omitempty"`
}

func (m *Validate) Name() string         { return "SieveScript/validate" }
func (m *Validate) Requires() []jmap.URI { return []jmap.URI{sieve.URI} }

type ValidateResponse struct {
	Account jmap.ID       `json:"accountId,omitempty"`
	Error   *jmap.SetError `json:"error"`
}

func newValidateResponse() jmap.MethodResponse { return &ValidateResponse{} }
