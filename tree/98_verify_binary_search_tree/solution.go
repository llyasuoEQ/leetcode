package verify_binary_search_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用中序遍历，然后看遍历的结果是不是一个有序的数组
// 时间复杂度：O(N+N) = 0(2N)
func isValidBST(root *TreeNode) bool {
	res := true
	if root == nil {
		return res
	}
	traversalRes := inOrderTraversal(root)
	tem := traversalRes[0]
	for i := 1; i < len(traversalRes); i++ {
		if traversalRes[i] < tem {
			res = false
			break
		}
		tem = traversalRes[i]
	}

	return res
}

func inOrderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inOrderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrderTraversal(root.Right)...)
	return res
}

// TODO 能否在前序遍历的同时就判断出是否平衡二叉树
// 利用搜索二叉树的性质去判断
func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}

func helper(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= high {
		return false
	}
	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
}
