package main

import "github.com/wzrtzt/GameServer/network"

func main() {
	server := network.NewServer(":9685")
	server.Start()
}
