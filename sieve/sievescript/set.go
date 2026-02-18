package sievescript

import (
	"github.com/mikluko/jmap"
	"github.com/mikluko/jmap/sieve"
)

type Set struct {
	Account   jmap.ID                    `json:"accountId,omitempty"`
	IfInState string                     `json:"ifInState,omitempty"`
	Create    map[jmap.ID]*SieveScript   `json:"create,omitempty"`
	Update    map[jmap.ID]jmap.Patch     `json:"update,omitempty"`
	Destroy   []jmap.ID                  `json:"destroy,omitempty"`

	// OnSuccessActivateScript is the ID of the script to activate if the
	// set operation succeeds, or "#creation-id" to reference a script being
	// created in the same call
	OnSuccessActivateScript *jmap.ID `json:"onSuccessActivateScript,omitempty"`

	// OnSuccessDeactivateScript indicates whether to deactivate the active
	// script if the set operation succeeds
	OnSuccessDeactivateScript bool `json:"onSuccessDeactivateScript,omitempty"`
}

func (m *Set) Name() string         { return "SieveScript/set" }
func (m *Set) Requires() []jmap.URI { return []jmap.URI{sieve.URI} }

type SetResponse struct {
	Account      jmap.ID                     `json:"accountId,omitempty"`
	OldState     string                      `json:"oldState,omitempty"`
	NewState     string                      `json:"newState,omitempty"`
	Created      map[jmap.ID]*SieveScript    `json:"created,omitempty"`
	Updated      map[jmap.ID]*SieveScript    `json:"updated,omitempty"`
	Destroyed    []jmap.ID                   `json:"destroyed,omitempty"`
	NotCreated   map[jmap.ID]*jmap.SetError  `json:"notCreated,omitempty"`
	NotUpdated   map[jmap.ID]*jmap.SetError  `json:"notUpdated,omitempty"`
	NotDestroyed map[jmap.ID]*jmap.SetError  `json:"notDestroyed,omitempty"`
}

func newSetResponse() jmap.MethodResponse { return &SetResponse{} }
