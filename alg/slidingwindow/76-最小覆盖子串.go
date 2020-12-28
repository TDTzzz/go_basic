package main

import (
	"log"
	"math"
)

func main() {
	//s := "ADOBECODEBANC"
	//t := "ABC"
	s := "aa"
	t := "aa"

	//log.Println(s[2:5])
	res := minWindow(s, t)
	log.Println(res)
}

func minWindow(s, t string) string {
	left, right, valid := 0, 0, 0
	need := make(map[string]int)
	window := make(map[string]int)

	start := 0
	l := math.MaxUint32
	for _, v := range t {
		need[string(v)]++
	}
	for right < len(s) {
		tmp := string(s[right])
		//判断是否是t里的
		right++
		if need[tmp] > 0 {
			window[tmp]++
			if window[tmp] == need[tmp] {
				valid++

			}
		}
		for valid == len(need) {
			//收紧need
			if right-left < l {
				start = left
				l = right - left
			}
			d := string(s[left])
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if l == math.MaxUint32 {
		return ""
	}
	return s[start : l+start]
}
