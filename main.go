package main

import (
	"log"
	"path/filepath"
)

func main() {
	join := filepath.Join("main.go")
	log.Println(join)
	dir := filepath.Dir(join)
	log.Println(dir)

	res, _ := filepath.Abs("main.go")
	log.Println(res)
}
