package blob

import "sarak.sh/mikluko/jmap"

func init() {
	jmap.RegisterMethod("Blob/copy", newCopyResponse)
}
