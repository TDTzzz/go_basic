package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial err", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	log.Println(reply)
}
