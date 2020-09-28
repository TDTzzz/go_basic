package main

import (
	"fmt"
	"hash/crc32"
	"log"
)

func main() {

	res := fmt.Sprintf("%s/%d", "svc.user.agent", crc32.ChecksumIEEE([]byte("127.0.0.1:8881")))
	log.Println(res)
}
