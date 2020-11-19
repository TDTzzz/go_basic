package main

import "fmt"

func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	key := n / 2
	left := MergeSort(arr[0:key])
	right := MergeSort(arr[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	newArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}
		} else {
			newArr[index] = left[i]
			i++
			index++
			if i == len(left) {
				copy(newArr[index:], right[j:])
			}
		}
	}
	return newArr
}

func main() {
	array := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	fmt.Println(array)
	array = MergeSort(array)
	fmt.Println(array)
}
