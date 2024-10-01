package tagmanager_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func dump(t *testing.T, i interface{ MarshalJSON() ([]byte, error) }) {
	t.Helper()
	var ret bytes.Buffer
	out, err := i.MarshalJSON()
	require.NoError(t, err)
	require.NoError(t, json.Indent(&ret, out, "", "  "))
	t.Log(ret.String())
}
