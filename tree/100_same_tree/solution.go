package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 利用深度优先遍历实现
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	// 先看当前节点是否相同
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// 当前节点的子树是否相同
	return p.Val == q.Val && isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
}

// 方法二：
// 利用广度优先遍历实现
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	pp, qq := []*TreeNode{p}, []*TreeNode{q}
	for len(pp) > 0 && len(qq) > 0 {
		node1, node2 := pp[0], qq[0]
		pp, qq = pp[1:], qq[1:]
		if node1.Val != node2.Val {
			return false
		}
		if (node1.Left == nil && node2.Left != nil) || (node1.Left != nil && node2.Left == nil) {
			return false
		}
		if (node1.Right == nil && node2.Right != nil) || (node1.Right != nil && node2.Right == nil) {
			return false
		}
		if node1.Left != nil {
			pp = append(pp, node1.Left)
		}
		if node1.Right != nil {
			pp = append(pp, node1.Right)
		}
		if node2.Left != nil {
			qq = append(qq, node2.Left)
		}
		if node2.Right != nil {
			qq = append(qq, node2.Right)
		}
	}
	return len(pp) == 0 && len(qq) == 0
}
