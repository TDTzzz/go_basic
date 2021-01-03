package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "0.0.0.0:8000")
	if err != nil {
		fmt.Printf("error %v connecting!", err)
	}
	conn.Write([]byte("sada"))
}
