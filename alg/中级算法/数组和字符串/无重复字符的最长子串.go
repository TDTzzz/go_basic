package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		ans = max(ans, rk+1-i)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}
