package peer_store

import (
	"github.com/wchen99998/torrent/metainfo"

	"github.com/wchen99998/dht/v2/krpc"
)

type InfoHash = metainfo.Hash

type Interface interface {
	AddPeer(InfoHash, krpc.NodeAddr)
	GetPeers(InfoHash) []krpc.NodeAddr
}
