package level_order_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路是和利用队列实现是一样的
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	m := []*TreeNode{root}
	for i := 0; len(m) > 0; i++ {
		res = append(res, []int{})
		n := []*TreeNode{}
		for j := 0; j < len(m); j++ {
			node := m[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				n = append(n, node.Left)
			}
			if node.Right != nil {
				n = append(n, node.Right)
			}
		}
		m = n
	}
	return res
}
