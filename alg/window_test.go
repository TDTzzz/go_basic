package main

//滑动窗口相关
func main() {

}

//最长无重复字符串
func lengthOfLongestSubstring(s string) (res int) {
	window := make(map[byte]int)
	left, right := 0, 0

	for i := 0; i < len(s); i++ {
		window[s[i]]++
		right++
		for window[s[i]] > 1 {
			window[s[left]]--
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}



