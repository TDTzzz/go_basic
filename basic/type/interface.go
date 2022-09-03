package main

import "fmt"

type Coder interface {
	code()
}

type Gopher struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

func main() {

	var aa *MyError
	fmt.Printf("err:%T,v:%v", aa, aa)
	return
	err := process()
	fmt.Printf("err:%T,v:%v", err, err)
	fmt.Println(err == nil)
	return
	//
	var c Coder
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	var g *Gopher
	fmt.Printf("g:%T,v:%v", g, g)
	fmt.Println(g == nil)

	//
	fmt.Printf("c: %T, %v\n", c, c)
	c = g
	fmt.Println(c == nil) //FALSE，因为c=g让c的动态类型变了
	fmt.Printf("c: %T, %v\n", c, c)
}

type MyError struct {
}

func (i MyError) Error() string {
	return "error"
}

func process() error {
	var err *MyError
	fmt.Printf("c: %T, %v\n", err, err)
	fmt.Println(err == nil)
	return err
}
