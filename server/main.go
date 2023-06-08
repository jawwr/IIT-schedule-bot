package main

import (
	"server/mainServer"
	_ "server/mainServer"
)

func main() {
	server := mainServer.New(8081)
	server.Start()
}
