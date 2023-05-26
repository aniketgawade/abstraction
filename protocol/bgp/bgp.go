package bgp

import (
	"fmt"

	"example.com/abstraction/protocol"
)

type bgp struct {
	protocolName string
}

func NewBgp() protocol.Protocol {
	return &bgp{
		protocolName: "bgp",
	}
}
func (b *bgp) Start() {
	fmt.Println("starting ", b.protocolName)
}

func (b *bgp) Stop() {
	fmt.Println("stopping ", b.protocolName)
}
