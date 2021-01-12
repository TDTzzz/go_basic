package main

import "fmt"

//逃逸分析
func main() {
	//fmt.Println("Called stackAnalysis", stackAnalysis())
	fmt.Println("Called heapAnalysis", heapAnalysis())
}

//go:noinline
//func stackAnalysis() int {
//	data := 55
//	return data
//}

//go:noinline
func heapAnalysis() *int {
	data := 55
	return &data
}
