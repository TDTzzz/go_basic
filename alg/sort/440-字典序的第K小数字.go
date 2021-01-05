package main

import "math"

func findKthNumber(n int, k int) int {

	// 此题是十叉树，第一层为1-9，第2层开始为0-9的十叉树
	p := 1      // 游标，指当前第p个元素
	prefix := 1 // 当前前缀节点（元素）

	for p < k {

		count := 0
		currentNode := prefix
		nextNode := prefix + 1

		// 计算当前节点currentNode的子节点个数count
		for currentNode <= n {
			count += int(math.Min(float64(nextNode), float64(n+1))) - currentNode
			currentNode *= 10
			nextNode *= 10
		}

		// 若游标所在位置，即p个元素，加上当前节点currentNode的子节点数之和，大于k，说明需查找的元素就在当前节点下；
		// 乘以10，即为往下一层查找
		// 游标移动到下一个元素
		if p+count > k {
			prefix *= 10
			p++
		} else {
			// 反之；说明要找的第k小元素不在当前节点下
			// 需要往后移动一个节点作为前缀，即前缀值+1
			// 游标跳过当前节点的所有子节点
			prefix++
			p += count
		}
	}

	// 最终游标p==k，即找到了第k小元素；而经过移动的前缀值prefix值，就是我们要找的元素值
	return prefix
}
