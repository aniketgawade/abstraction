package main

import (
	"fmt"

	"example.com/abstraction/protocol"
	"example.com/abstraction/protocol/bgp"
	"example.com/abstraction/protocol/ospf"
)

func main() {
	processIntegrationObj("ospf")
	processIntegrationObj("bgp")
}

func processIntegrationObj(blobType string) {
	var protocol protocol.Protocol
	switch blobType {
	case "ospf":
		protocol = ospf.NewOSPF()
	case "bgp":
		protocol = bgp.NewBgp()
	default:
		fmt.Println("unsupported")
	}

	protocol.Start()
	protocol.Stop()
}
