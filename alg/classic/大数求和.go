package main

import (
	"strconv"
	"strings"
)

//字符串翻转
func reverse(str string) string {
	tmp := []byte(str)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}

//大数相加
func addString(str1, str2 string) string {
	tmp := []string{}
	flag := 0

	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0 || flag != 0; i, j = i-1, j-1 {
		x1, x2 := 0, 0
		if i < 0 {
			x1 = 0
		} else {
			x1 = int(str1[i] - '0')
		}

		if j < 0 {
			x2 = 0
		} else {
			x2 = int(str2[j] - '0')
		}

		sum := (x1 + x2 + flag) % 10
		sums := strconv.Itoa(sum)
		tmp = append(tmp, sums)
		flag = (x1 + x2 + flag) / 10
	}

	return reverse(strings.Join(tmp, ""))
}

//大数相乘
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var res string

	for i := len(num2) - 1; i >= 0; i-- {
		if num2[i] == '0' {
			continue
		}

		tmp := []string{}
		for j := 0; j < len(num2)-1-i; j++ {
			tmp = append(tmp, "0")
		}
		x2 := int(num2[i] - '0')
		flag := 0
		for n := len(num1) - 1; n >= 0 || flag != 0; n-- {
			var x1 int
			if n < 0 {
				x1 = 0
			} else {
				x1 = int(num1[n] - '0')
			}

			mul := (x1*x2 + flag) % 10
			muls := strconv.Itoa(mul)
			tmp = append(tmp, muls)
			flag = (x1*x2 + flag) / 10
		}

		res = addString(res, reverse(strings.Join(tmp, "")))
	}
	return res
}
