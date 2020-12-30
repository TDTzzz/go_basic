package main

func lengthOfLongestSubstring(s string) (res int) {
	window := make(map[string]int)
	left, right := 0, 0

	for right < len(s) {
		c := string(s[right])
		right++
		window[c]++
		for window[c] > 1 {
			d := string(s[left])
			left++
			window[d]--
		}
		if res < right-left {
			res = right - left
		}
	}
	return
}
