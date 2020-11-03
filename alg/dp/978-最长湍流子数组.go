package main

import "log"

//超级烂的代码（dp思路）
func maxTurbulenceSize(A []int) int {
	if len(A) < 2 {
		return len(A)
	}
	dp := make([]int, len(A))
	dp2 := make([]int, len(A))
	//初始化每个数都为1
	for k, _ := range dp {
		dp[k] = 1
		dp2[k] = 1
	}

	res := 0
	for i := 1; i < len(A); i++ {
		cur, prev := A[i], A[i-1]
		if i%2 == 0 {
			if cur > prev {
				dp[i] = dp[i-1] + 1

			} else {
				dp[i] = 1
			}
		} else {
			if cur < prev {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = 1
			}
		}
		if i%2 == 1 {
			if cur > prev {
				dp2[i] = dp2[i-1] + 1

			} else {
				dp2[i] = 1
			}
		} else {
			if cur < prev {
				dp2[i] = dp2[i-1] + 1
			} else {
				dp2[i] = 1
			}
		}
	}
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	for _, v := range dp2 {
		if v > res {
			res = v
		}
	}
	return res
}

//好一点的DP
func maxTurbulenceSizeV2(A []int) int {
	dp_1, dp_2, ans := 1, 1, 1
	for i := 1; i < len(A); i++ {
		if cur, prev := A[i], A[i-1]; cur > prev {
			dp_1, dp_2 = 1, dp_1+1
			if dp_2 > ans {
				ans = dp_2
			}
		} else if cur == prev {
			dp_1, dp_2 = 1, 1
		} else {
			dp_1, dp_2 = dp_2+1, 1
			if dp_1 > ans {
				ans = dp_1
			}
		}
	}
	return ans
}

//好思路：滑动窗口
func maxTurbulenceSizeV3(A []int) int {
	ans, i, legt := 1, 1, len(A) //i遍历A时的索引
	for {
		for i < legt && A[i] == A[i-1] { //先把一连串的==全都过掉 比如[1,1,1,2]直接来到i==3的位置
			i++
		}
		if i == legt {
			return ans
		}
		l, flag := 2, A[i] > A[i-1] //True:↗ False:↘ l表示当前的湍流子数组的长度
		i++
		for i < legt && A[i] != A[i-1] && (A[i] > A[i-1]) != flag {
			l++
			i++
			flag = !flag //转换
		}
		if l > ans {
			ans = l
		}
	}
}

func main() {
	A := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	res := maxTurbulenceSize(A)
	log.Println(res)
}
