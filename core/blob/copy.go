package blob

import (
	"sarak.sh/mikluko/jmap"
)

// Copy a binary blob from one account to another
// https://www.rfc-editor.org/rfc/rfc8620.html#section-6.3
type Copy struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"blobIds,omitempty"`
}

func (m *Copy) Name() string { return "Blob/copy" }

func (m *Copy) Requires() []jmap.URI { return nil }

// The response to a Blob/copy call
// https://www.rfc-editor.org/rfc/rfc8620.html#section-6.3
type CopyResponse struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	Copied map[jmap.ID]jmap.ID `json:"copied,omitempty"`

	NotCopied map[jmap.ID]*jmap.SetError `json:"notCopied,omitempty"`
}

func newCopyResponse() jmap.MethodResponse { return &CopyResponse{} }
