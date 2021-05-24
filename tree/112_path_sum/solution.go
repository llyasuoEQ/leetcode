package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现 - 深度优先遍历
// 时间复杂度：O(N)
// 空间复杂度：O(H)
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
}

// 迭代实现 - 广度优先遍历
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	qNode := []*TreeNode{root}
	qSum := []int{root.Val}
	for len(qNode) > 0 {
		nowNode := qNode[0]
		nowSum := qSum[0]
		qNode = qNode[1:]
		qSum = qSum[1:]
		if nowNode.Left == nil && nowNode.Right == nil {
			if nowSum == targetSum {
				return true
			}
			continue
		}
		if nowNode.Left != nil {
			qNode = append(qNode, nowNode.Left)
			qSum = append(qSum, nowSum+nowNode.Left.Val)
		}
		if nowNode.Right != nil {
			qNode = append(qNode, nowNode.Right)
			qSum = append(qSum, nowSum+nowNode.Right.Val)
		}
	}
	return false
}
