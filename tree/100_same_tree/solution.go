package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 利用深度优先遍历实现
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 先看当前节点是否相同
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// 当前节点的子树是否相同
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 方法二：
// 利用广度优先遍历实现
