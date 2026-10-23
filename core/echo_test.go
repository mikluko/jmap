package core

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sarak.sh/mikluko/jmap"
)

func TestEcho(t *testing.T) {
	echo := &Echo{
		Hello: "world",
	}

	req := jmap.Request{}
	req.Invoke(echo)
	data, err := json.Marshal(req)
	assert.NoError(t, err)
	exp := `{"using":[],"methodCalls":[["Core/echo",{"Hello":"world"},"0"]]}`
	assert.Equal(t, exp, string(data))
}
