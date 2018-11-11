# Custom DHT Servers

https://github.com/anacrolix/torrent/search?p=1&q=dht&unscoped_q=dht

## Types

* dht.Server
* dht.ServerConfig
* dht.ServerStats
* dht.Peer
* dht.StartingNodesGetter
* dht.Addr
* dht/krpc.NodeInfo
* dht/krpc.NodeAddr
* torrent/metainfo.Hash

## config.go

* type ClientConfig struct
  - dht.StartingNodesGetter

* type NewDefaultConfig struct
  - dht.GlobalBootstrapAddrs

## client.go

* type Client struct
  - dht.Server

* func NewClient(cfg *ClientConfig) (cl *Client, err error)
  - dht.Server

* func writeDhtServerStatus(w io.Writer, s *dht.Server)
  - dht.Server.Stats()
  - dht.ServerStats
  - dht.ServerStats.Nodes
  - dht.ServerStats.GoodNodes
  - dht.ServerStats.BadNodes
  - dht.ServerStats.SuccessfulOutboundAnnouncePeerQueries
  - dht.ServerStats.OutstandingTransactions
  - dht.Server.ID()

* func (cl *Client) WriteStatus(_w io.Writer)
  - dht.Server

* func (cl *Client) newDhtServer(conn net.PacketConn) (s *dht.Server, err error)
  - dht.ServerConfig
  - dht.NewServer()
  - dht.Server
  - dht.Server.Bootstrap()

* func (cl *Client) eachDhtServer(f func(*dht.Server))
  - dht.Server

* func (cl *Client) Close()
  - dht.Server

* func (cl *Client) dhtPort() (ret uint16)
  - dht.Server
  - dht.Server.Addr()

* func (cl *Client) haveDhtServer() (ret bool)
  - dht.Server

* func (cl *Client) AddTorrentInfoHashWithStorage(infoHash metainfo.Hash, specStorage storage.ClientImpl) (t *Torrent, new bool)
  - dht.Server

* func (cl *Client) onDHTAnnouncePeer(ih metainfo.Hash, p dht.Peer)
  - dht.Peer
  - peer.IP
  - peer.Port

* func (cl *Client) AddDHTNodes(nodes []string)
  - dht/krpc.NodeInfo
  - dht/krpc.NodeAddr
  - dht.Server
  - dht.Server.AddNode()

* func (cl *Client) DhtServers() []*dht.Server
  - dht.server

## torrent.go

* func (t *Torrent) consumeDHTAnnounce(pvs -chan dht.PeersValues)
 - dht.PeersValues

* func (t *Torrent) announceDHT(impliedPort bool, s *dht.Server) (err error)
  - dht.Server
  - dht.Announce
  - dht.PeersValues
  - dht.Server.Announce()

* func (t *Torrent) dhtAnnouncer(s *dht.Server)
  - dht.Server

## connection.go

* func (c *connection) mainReadLoop() (err error)
  - dht.Server
  - dht.Server.Ping()

## Peer.go

* func (me *Peer) FromPex(na krpc.NodeAddr, fs peer_protocol.PexPeerFlags)
  - krpc.NodeAddr

## Peers.go

* func (me *Peers) AppendFromPex(nas []krpc.NodeAddr, fs []peer_protocol.PexPeerFlags)
  - krpc.NodeAddr

## peer_protocol/pex.go

* type PexMsg struct
  - krpc.CompactIPv4NodeAddrs()
  - krpc.CompactIPv6NodeAddrs()

## tracker.go

* type Announce struct
  - krpc.NodeAddr

## tracker_scraper.go

* func (me *trackerScraper) announce() (ret trackerAnnounceResult)
  - krpc.NodeAddr

## tracker/udp.go

* func (c *udpAnnounce) Do(req AnnounceRequest) (res AnnounceResponse, err error)
  - krpc.NodeAddr
  - krpc.CompactIPv4NodeAddrs
  - krpc.CompactIPv6NodeAddrs

## tracker/http.go

* type httpResponse struct
  - krpc.CompactIPv6NodeAddrs

* func (me *Peers) UnmarshalBencode(b []byte) (err error)
  - krpc.CompactIPv4NodeAddrs

## tracker/peer.go

* func (p Peer) FromNodeAddr(na krpc.NodeAddr) Peer
  - krpc.NodeAddr

## tracker/server.go

* type torrent struct
  - krpc.NodeAddr

* func (s *server) serveOne() (err error)
  - krpc.CompactIPv4NodeAddrs()
  - krpc.CompactIPv6NodeAddrs()
