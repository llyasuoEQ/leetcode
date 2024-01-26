package minimum_depth_of_binary_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 递归实现-深度优先遍历啊
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth1(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth1(root.Right), minD)
	}

	return minD + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 方法二：
// 迭代实现-广度优先遍历
// 时间复杂度：O（N）
// 空间复杂度：0（N）
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minD := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		minD++
		queueTemp := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return minD
			}
			if node.Left != nil {
				queueTemp = append(queueTemp, node.Left)
			}
			if node.Right != nil {
				queueTemp = append(queueTemp, node.Right)
			}
		}
		queue = queueTemp
	}
	return minD
}
