package krpc

import (
	"github.com/wchen99998/torrent/metainfo"
)

type Bep46Payload struct {
	Ih metainfo.Hash `bencode:"ih"`
}
