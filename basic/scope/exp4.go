package main

import "fmt"

func main() {
	p := Person{"张三"}
	modify(p)
	fmt.Println(p)
	modify2(&p)
	fmt.Println(p)
}

type Person struct {
	Name string
}

func modify(p Person) {
	p.Name = "李四" //不改原值
}

//其实指针类型的参数也是值传递
func modify2(p *Person) {
	p.Name = "李四" //改原值
}
