package krpc

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wchen99998/torrent/bencode"
)

// https://github.com/wchen99998/torrent/issues/166
func TestUnmarshalBadError(t *testing.T) {
	var e Error
	err := bencode.Unmarshal([]byte(`l5:helloe`), &e)
	require.Error(t, err)
}
