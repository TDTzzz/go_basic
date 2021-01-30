package main

//main

func main() {

}

//大整数
func twoSum(s1, s2 string) string {

	l1, l2 := len(s1), len(s2)

	var lowL int
	var res string
	//var  string
	if l1 < l2 {
		lowL = l1

	} else {
		lowL = l2
	}

	tmp := 0
	for i := 0; i < lowL; i++ {
		currNum := int(s1[i])
		currNum2 := int(s2[i])

		currSum := currNum + currNum2 + tmp
		if currSum > 10 {
			tmp = 1
			currSum -= 10
		} else {
			tmp = 0
		}
		//字符串拼接

	}

	//补位处理



	return res
}
