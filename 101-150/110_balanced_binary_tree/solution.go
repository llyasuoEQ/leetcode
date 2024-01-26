package balanced_binary_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lHeight := height(root.Left)
	rHeight := height(root.Right)

	if math.Abs(float64(lHeight-rHeight)) < 2 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := height(root.Left) + 1
	rHeight := height(root.Right) + 1

	if lHeight > rHeight {
		return lHeight
	}
	return rHeight
}

func height1(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	lHeight := height1(root.Left) + 1
	rHeight := height1(root.Right) + 1
	return math.Max(lHeight, rHeight)
}
