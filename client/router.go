package client

import (
	"log"
	"net"

	"github.com/songgao/water/waterutil"
)

const (
	// DestinationLocal defines the local routing of a packet.
	DestinationLocal = iota
	// DestinationRemote defines the remote routing of a packet.
	DestinationRemote
	// SourceLocal stands for a local packet origin.
	SourceLocal = iota
	// SourceRemote stands for a remote packet origin.
	SourceRemote
)

// A Packet wraps the payload, source or destination.
type Packet struct {
	Data    []byte
	SrcType int
	DstType int
	DstAddr string
}

// A Router is an instance of an inbound and outbound channel with filtering.
type Router struct {
	inbound  chan Packet
	outbound chan Packet
	localIPs []net.IP
	subnets  []*net.IPNet
}

// NewRouter creates and returns a Router instance.
// It receives the in- and outbound channel as well
// as a list of local IPs and subnets for filtering packets.
func NewRouter(inbound, outbound chan Packet, localIPs []net.IP, subnets []*net.IPNet) *Router {
	return &Router{
		inbound:  inbound,
		outbound: outbound,
		localIPs: localIPs,
		subnets:  subnets,
	}
}

// Route starts the main routing loop.
func (r *Router) Route() error {
	for {
		select {
		case pkt := <-r.inbound:
			dstIP := waterutil.IPv4Destination(pkt.Data)

			// Filtering packets that are not intended for the assigned subnet.
			if !r.isForSubnet(dstIP) {
				log.Printf("Skipping packet to %s", dstIP.String())
				continue
			}

			// Routing local packet to local interface.
			if r.isLocal(dstIP) {
				r.outbound <- Packet{Data: pkt.Data, DstType: DestinationLocal}
				continue
			}

			r.outbound <- Packet{
				Data:    pkt.Data,
				DstType: DestinationRemote,
				DstAddr: dstIP.String(),
			}
		}
	}
}

func (r *Router) isForSubnet(dstIP net.IP) bool {
	for _, s := range r.subnets {
		if s.Contains(dstIP) {
			return true
		}
	}
	return false
}

func (r *Router) isLocal(dstIP net.IP) bool {
	for _, ip := range r.localIPs {
		if ip.Equal(dstIP) {
			return true
		}
	}
	return false
}
