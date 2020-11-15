package main

func minHeightShelves(books [][]int, shelf_width int) int {
	dp := make([]int, len(books)+1)
	//dp[i]表示以第i本书作为一层的结尾所最后得到的总高度
	for i := 1; i <= len(books); i++ {
		tempwidth := books[i-1][0]
		tempheight := books[i-1][1]
		dp[i] = tempheight + dp[i-1]
		for j := i - 1; j > 0; j-- {
			if tempwidth+books[j-1][0] <= shelf_width {
				tempwidth += books[j-1][0]
				tempheight = max(tempheight, books[j-1][1])
				dp[i] = min(dp[i], dp[j-1]+tempheight)
			} else {
				break
			}
		}
	}
	return dp[len(books)]
}
func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
