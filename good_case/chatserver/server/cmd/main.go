package main

import "go_basic/good_case/chatserver/server"

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	// start the server
	s.Start()
}
