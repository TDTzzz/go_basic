package main

import (
	"fmt"
	"log"
)

type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func main() {
	add := address{province: "北京", city: "北京"}
	printString(add)
	printString(&add)

	var si fmt.Stringer = address{province: "上海", city: "上海"}
	printString(si)
	//sip:=&si
	//printString(sip)

	//更改值
	p := person{
		name: "bb",
		age:  18,
	}
	modifyPerson(&p)
	log.Println(p)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

type person struct {
	name string
	age  int
}

//更换值
func modifyPerson(p *person) {
	p.name = "啊啊"
	p.age = 20
}
