package main

import (
	"fmt"
	"math"
)

func main() {

	res := divide(15, 2)
	fmt.Println(res)
}

//笨蛋减法
func divide(a, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	res := 0
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	for a >= b {
		a -= b
		res++
	}
	return sign * res
}

//聪明一点的办法（成倍减）

// 超时 时间复杂度：O(logn * logn)，n 是最大值 2147483647 --> 10^10
func divide2(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}

	res := 0
	for a <= b {
		value, k := b, 1
		// 0xc0000000 是十进制 -2^30 的十六进制的表示
		// 判断 value >= 0xc0000000 的原因：保证 value + value 不会溢出
		// 可以这样判断的原因是：0xc0000000 是最小值 -2^31 的一半，
		// 而 a 的值不可能比 -2^31 还要小，所以 value 不可能比 0xc0000000 小
		// -2^31 / 2 = -2^30
		for value >= 0xc0000000 && a <= value+value {
			value += value
			k += k
		}
		a -= value
		res += k
	}
	return sign * res
}

//位运算

// 时间复杂度：O(1)
func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	a = abs(a)
	b = abs(b)

	res := 0
	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			a = a - (b << i)
			res += 1 << i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
