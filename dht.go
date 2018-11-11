package torrent

import (
	"fmt"
	"net"
	"strconv"
)

// DHT is the torrent compatible interface dht servers must implement.
type DHT interface {
	Addr() net.Addr
	Stats() ServerStats
	ID() [20]byte
	Close()
	Announce(infoHash [20]byte, port int, impliedPort bool) (<-chan PeersValues, error)
	AddNode(IP net.IP, port int) error
}

type NodeAddr struct {
	IP   net.IP
	Port int
}

type NodeInfo struct {
	ID   [20]byte
	Addr NodeAddr
}

type PeersValues struct {
	Peers    []NodeAddr // Peers given in get_peers response.
	NodeInfo            // The node that gave the response.
}

func (me NodeInfo) String() string {
	return fmt.Sprintf("{%x at %s}", me.ID, me.Addr)
}

// A zero Port is taken to mean no port provided, per BEP 7.
func (me NodeAddr) String() string {
	if me.Port == 0 {
		return me.IP.String()
	}
	return net.JoinHostPort(me.IP.String(), strconv.FormatInt(int64(me.Port), 10))
}

// ServerStats instance is returned by DHT.Stats() and stores Server metrics
type ServerStats struct {
	// Count of nodes in the node table that responded to our last query or
	// haven't yet been queried.
	GoodNodes int
	// Count of nodes in the node table.
	Nodes int
	// Transactions awaiting a response.
	OutstandingTransactions int
	// Individual announce_peer requests that got a success response.
	SuccessfulOutboundAnnouncePeerQueries int64
	// Nodes that have been blocked.
	BadNodes                 uint
	OutboundQueriesAttempted int64
}
