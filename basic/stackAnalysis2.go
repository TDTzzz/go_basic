package main

type smallStruct struct {
	a, b int64
	c, d float64
}

func main() {
	smallAllocation()
}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}
