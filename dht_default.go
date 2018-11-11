package torrent

import (
	"log"
	"net"

	"github.com/anacrolix/dht"
	"github.com/anacrolix/dht/krpc"
	"github.com/anacrolix/torrent/metainfo"
)

type DefaultDHT struct {
	*dht.Server
}

func (me *DefaultDHT) Stats() ServerStats {
	return ServerStats(me.Server.Stats())
}

func (me *DefaultDHT) AddNode(ip net.IP, port int) error {
	ni := krpc.NodeInfo{
		Addr: krpc.NodeAddr{
			IP:   ip,
			Port: port,
		},
	}
	return me.Server.AddNode(ni)
}

func (me *DefaultDHT) Announce(infoHash [20]byte, port int, impliedPort bool) (<-chan PeersValues, error) {
	a, err := me.Server.Announce(infoHash, port, impliedPort)
	if err != nil {
		return nil, err
	}
	pvs := make(chan PeersValues, 0)
	go func() {
		for v := range a.Peers {
			peers := make([]NodeAddr, len(v.Peers))
			for _, p := range v.Peers {
				peers = append(peers, NodeAddr(p))
			}
			pv := PeersValues{
				NodeInfo: NodeInfo{
					ID:   v.ID,
					Addr: NodeAddr(v.Addr),
				},
				Peers: peers,
			}
			pvs <- pv
		}
	}()
	return pvs, nil
}

func DefaultDhtServer(cl *Client, conn net.PacketConn) (s DHT, err error) {
	onAnnounce := func(ih metainfo.Hash, p dht.Peer) {
		cl.onDHTAnnouncePeer(ih, p.IP, p.Port)
	}
	cfg := dht.ServerConfig{
		IPBlocklist:    cl.ipBlockList,
		Conn:           conn,
		OnAnnouncePeer: onAnnounce,
		PublicIP: func() net.IP {
			if connIsIpv6(conn) && cl.config.PublicIp6 != nil {
				return cl.config.PublicIp6
			}
			return cl.config.PublicIp4
		}(),
		StartingNodes: dht.GlobalBootstrapAddrs,
	}
	ds, err := dht.NewServer(&cfg)
	if err == nil {
		go func() {
			if _, err := ds.Bootstrap(); err != nil {
				log.Printf("error bootstrapping dht: %s", err)
			}
		}()
	}
	return &DefaultDHT{ds}, err
}
