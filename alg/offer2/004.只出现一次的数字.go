package main

//笨办法哈希表
func singleNumber(nums []int) int {

	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for num, occ := range freq {
		if occ == 1 {
			return num
		}
	}
	return 0
}

//依次确定每一个二进制位
func singleNumber2(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}

	return int(ans)
}

//数字电路的设计

func singleNumber3(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a, b = b&^a&num|a&^b&^num, (b^num)&^a
	}
	return b
}
