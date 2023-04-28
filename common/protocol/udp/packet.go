package udp

import (
	"github.com/gptlocal/netools/common/buf"
	"github.com/gptlocal/netools/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
