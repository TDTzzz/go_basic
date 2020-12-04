package offer

//滑动窗口+hash
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 存储 left~right 中间的字符
	set := make(map[byte]bool)

	left := 0
	set[s[0]] = true
	maxLen := 1
	for right := 1; right < len(s); {
		b := set[s[right]]
		if !b {
			set[s[right]] = true
			right++
		} else {
			delete(set, s[left])
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}

//V2:动态规划+hash

