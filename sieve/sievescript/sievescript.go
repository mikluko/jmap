package sievescript

import "github.com/mikluko/jmap"

func init() {
	jmap.RegisterMethod("SieveScript/get", newGetResponse)
	jmap.RegisterMethod("SieveScript/set", newSetResponse)
	jmap.RegisterMethod("SieveScript/changes", newChangesResponse)
	jmap.RegisterMethod("SieveScript/query", newQueryResponse)
	jmap.RegisterMethod("SieveScript/queryChanges", newQueryChangesResponse)
	jmap.RegisterMethod("SieveScript/validate", newValidateResponse)
}

// SieveScript represents a Sieve script on the server
type SieveScript struct {
	// ID is the server-assigned identifier for this script
	ID jmap.ID `json:"id,omitempty"`

	// Name is the user-assigned name of the script
	Name *string `json:"name,omitempty"`

	// BlobID is the identifier for the blob containing the raw script
	BlobID jmap.ID `json:"blobId,omitempty"`

	// IsActive indicates whether this script is the active script
	IsActive bool `json:"isActive,omitempty"`
}
