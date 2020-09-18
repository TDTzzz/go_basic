package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string

	err = client.Call("HelloService.Hello", "haaaa", &reply)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}
