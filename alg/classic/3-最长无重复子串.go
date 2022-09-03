package main

func main() {

	s := "abcabcbb"

	res := lengthOfLongestSubstring(s)
}

//滑动窗口
func lengthOfLongestSubstring(s string) (res int) {
	sLen := len(s)
	right, m := -1, make(map[byte]int)

	for l := 0; l < sLen; l++ {
		if l != 0 {
			delete(m, s[l-1])
		}

		for right+1 < sLen && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		res = max(res, right-l+1)
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
