package blob

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sarak.sh/mikluko/jmap"
)

func TestCopyResponse(t *testing.T) {
	data := []byte(`{
		"fromAccountId": "A1",
		"accountId": "A2",
		"copied": {"B1": "B9"},
		"notCopied": {"B2": {"type": "notFound"}}
	}`)

	var resp CopyResponse
	assert.NoError(t, json.Unmarshal(data, &resp))
	assert.Equal(t, jmap.ID("A1"), resp.FromAccount)
	assert.Equal(t, jmap.ID("A2"), resp.Account)
	assert.Equal(t, map[jmap.ID]jmap.ID{"B1": "B9"}, resp.Copied)
	assert.Contains(t, resp.NotCopied, jmap.ID("B2"))
	assert.Equal(t, "notFound", resp.NotCopied["B2"].Type)
}
