package main

import (
	"github.com/tidwall/gjson"
	"log"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	gjsonTest()
}

func gjsonTest() {
	value := gjson.Parse(json).Map()
	log.Println(value["name"])
}
