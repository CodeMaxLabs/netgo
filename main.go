package main

import "netgo/pkg/network"

func main() {
	// pkg.NewUDPServer(8888)

	network.NewHttpWebsocketServer(23333, "/messages")
}
