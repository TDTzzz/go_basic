package main

import "log"

func main() {

	s := "helloworld"
	t := "oow"
	res := checkInclusion(s, t)
	log.Println(res)

}

func checkInclusion(s, t string) bool {
	need := make(map[string]int)
	window := make(map[string]int)

	//初始化need
	for _, v := range t {
		need[string(v)]++
	}

	left := 0
	right := 0
	valid := 0

	for right < len(s) {
		c := string(s[right])
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(t) {
			if valid == len(need) {
				return true
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
	return false
}
