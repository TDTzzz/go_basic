package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	heap(arr)
}

func heap(arr []int) {
	//构建大顶堆
	for i := len(arr) / 2; i >= 0; i-- {
		heap_help(arr, i, len(arr)-1)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heap_help(arr, 0, i-1)
	}

	log.Println(arr)
}

func heap_help(arr []int, dad int, end int) {
	son := 2*dad + 1
	if son > end {
		return
	}

	if son+1 <= end && arr[son] < arr[son+1] {
		son++
	}

	if arr[dad] < arr[son] {
		arr[dad], arr[son] = arr[son], arr[dad]
	}
	heap_help(arr, son, end)
}
