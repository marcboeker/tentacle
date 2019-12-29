package client

import (
	"github.com/songgao/water"
)

// Interface wraps a TUN interface.
type Interface struct {
	*water.Interface
}

// Read reads new packets from the interface and forwards them to the
// inbound channel.
func (i Interface) Read(inbound chan Packet) error {
	buf := make([]byte, 2000)
	for {
		n, err := i.Interface.Read(buf)
		if err != nil {
			return err
		}

		inbound <- Packet{Data: buf[:n], SrcType: SourceLocal}
	}
}
