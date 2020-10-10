package alg

//辅助函数

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x <= 0 {
		return -1 * x
	}
	return x
}
