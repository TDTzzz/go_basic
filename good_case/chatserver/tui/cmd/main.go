package main

import (
	"flag"
	"go_basic/good_case/chatserver/client"
	"go_basic/good_case/chatserver/tui"
	"log"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")

	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	tui.StartUi(client)
}
