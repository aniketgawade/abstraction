package ospf

import (
	"fmt"

	"example.com/abstraction/protocol"
)

type ospf struct {
	protocolName string
}

func NewOSPF() protocol.Protocol {
	return &ospf{
		protocolName: "ospf",
	}
}
func (b *ospf) Start() {
	fmt.Println("starting ", b.protocolName)
}

func (b *ospf) Stop() {
	fmt.Println("stopping ", b.protocolName)
}
