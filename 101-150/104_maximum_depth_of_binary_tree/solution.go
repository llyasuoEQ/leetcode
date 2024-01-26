package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法-：
// 递归实现
// 求最大深度
// 时间复杂度：O(N)
// 空间复杂度：O(Height)
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lH := maxDepth1(root.Left) + 1
	rH := maxDepth1(root.Right) + 1
	if lH > rH {
		return lH
	}
	return rH
}

// 方法二：
// 迭代实现
// 广度优先遍历（层序遍历） - 层数就是深度
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	currQ := []*TreeNode{root}
	i := 0
	for len(currQ) > 0 {
		nextQ := []*TreeNode{}
		for _, n := range currQ {
			if n.Left != nil {
				nextQ = append(nextQ, n.Left)
			}
			if n.Right != nil {
				nextQ = append(nextQ, n.Right)
			}
		}
		i++
		currQ = nextQ
	}
	return i
}
