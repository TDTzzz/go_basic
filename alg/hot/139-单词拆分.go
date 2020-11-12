package hot



func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}

	var hash = make(map[string]struct{})
	for _, word := range wordDict {
		hash[word] = struct{}{}
	}

	var dp = make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := hash[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

