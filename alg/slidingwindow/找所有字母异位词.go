package main

func findAnagrams(s, t string) (res []int) {
	need := make(map[string]int)
	window := make(map[string]int)

	for _, v := range t {
		need[string(v)]++
	}

	left := 0
	right := 0
	valid := 0

	for left < len(s) {
		c := string(s[right])
		right++
		if need[c] > 0 {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for right-left >= len(t) {
			if valid == len(need) {
				res = append(res, left)
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

	return
}
