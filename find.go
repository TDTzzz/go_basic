package main

import (
	"log"
	"sync"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	target := 7
	res := find(nums, target)
	log.Println(res)

	wg := sync.WaitGroup{

	}
}

func aa(*sync.WaitGroup) {

}

func find(nums []int, target int) bool {
	if len(nums) == 1 {
		return false
	}
	key := len(nums) / 2
	if target > nums[key] {

		return find(nums[key:], target)
	} else if target < nums[key] {
		return find(nums[:key], target)
	} else {
		return true
	}
}

type AA struct {
}

//id name
