package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1233")
	if err != nil {
		log.Fatal("dialog:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "Hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}
