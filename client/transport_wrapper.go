package client

import (
	"sync"

	"github.com/songgao/water/waterutil"
)

// TransportWrapper keeps track of all incoming and outgoing peer connections.
type TransportWrapper struct {
	rpc       *RPC
	transport *Transport
	pool      map[string]*Stream // Dest IP -> stream
	mutex     sync.Mutex
	inbound   chan Packet
}

// NewTransportWrapper returns an instance of a TransportWrapper.
func NewTransportWrapper(r *RPC, t *Transport, inbound chan Packet) *TransportWrapper {
	return &TransportWrapper{
		rpc:       r,
		transport: t,
		pool:      map[string]*Stream{},
		mutex:     sync.Mutex{},
		inbound:   inbound,
	}
}

// AddStream adds a new stream for the given peer address.
func (tw *TransportWrapper) AddStream(addr string, stream *Stream) {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()
	tw.pool[addr] = stream
}

// RemoveStream removes a stream for the given peer address.
func (tw *TransportWrapper) RemoveStream(addr string) {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()
	delete(tw.pool, addr)
}

// GetOrConnect checks if a stream for the given peer already exists.
// Otherwise it connects to the peer and returns the stream.
func (tw *TransportWrapper) GetOrConnect(addr string) (*Stream, error) {
	stream, ok := tw.pool[addr]
	if ok {
		if stream.Closed {
			tw.RemoveStream(addr)
		} else {
			return stream, nil
		}
	}

	peerAddr, err := tw.rpc.Connect(addr)
	if err != nil {
		return nil, err
	}

	stream, err = tw.transport.Dial(peerAddr, addr)
	if err != nil {
		return nil, err
	}

	tw.AddStream(addr, stream)

	// Monitor stream for incoming packets and write them to the inbound channel.
	go stream.Read(func(pkt []byte, s *Stream) {
		tw.inbound <- Packet{
			Data:    pkt,
			SrcType: SourceRemote,
		}
	})

	return stream, nil
}

// Serve handles incoming peer connections and
// forwards packets to the inbound channel.
func (tw *TransportWrapper) Serve(localAddr string) error {
	return tw.transport.Server(func(stream *Stream) {
		if err := stream.Read(func(pkt []byte, s *Stream) {
			srcIP := waterutil.IPv4Source(pkt)
			tw.AddStream(srcIP.String(), s)

			tw.inbound <- Packet{
				Data:    pkt,
				SrcType: SourceRemote,
			}
		}); err != nil {
			return
		}
	})
}
