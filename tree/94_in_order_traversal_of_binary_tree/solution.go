package in_order_traversal_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归的方式实现
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// TODO 遍历的方式
