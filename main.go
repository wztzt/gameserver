package main

import "github.com/wztzt/gameserver/network"

func main() {
	server := network.NewServer(":9685")
	server.Start()
}
