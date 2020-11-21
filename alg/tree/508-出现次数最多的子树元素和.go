package main

//本质是后序遍历
func findFrequentTreeSum(root *TreeNode) []int {
	res := make([]int, 0)
	d := make(map[int]int)
	var mostFreq int
	var findSum func(root *TreeNode) int
	findSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		//计算当前子树的和
		sum := root.Val + findSum(root.Left) + findSum(root.Right)
		d[sum]++
		if d[sum] > mostFreq {
			mostFreq = d[sum]
		}
		return sum
	}
	findSum(root)
	for k, v := range d {
		if v == mostFreq {
			res = append(res, k)
		}
	}
	return res
}
