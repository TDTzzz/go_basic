package main

func maxProduct2(words []string) int {
	// O(mn)
	var n = len(words)
	var masks = make([]int, n)
	for i := 0; i < n; i++ {
		var bitMask = 0
		for _, c := range words[i] {
			bitMask |= (1 << (c - 'a'))
		}
		masks[i] = bitMask
	}

	var ans = 0
	for i := range words {
		var word1 = words[i]
		for j := i + 1; j < len(words); j++ {
			var word2 = words[j]
			if (masks[i] & masks[j]) == 0 {
				var length = len(word1) * len(word2)
				if length > ans {
					ans = length
				}
			}
		}
	}
	return ans
}
