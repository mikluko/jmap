package blob

import "github.com/mikluko/jmap"

func init() {
	jmap.RegisterMethod("Blob/copy", newCopyResponse)
}
