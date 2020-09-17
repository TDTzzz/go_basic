package main

import (
	"go_basic/rpc/svc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("HelloService", new(svc.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("tcp err:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept err:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
