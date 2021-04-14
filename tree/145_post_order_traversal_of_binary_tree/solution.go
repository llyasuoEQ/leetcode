package post_order_traversal_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// TODO 进阶 利用迭代的方式实现
func postorderTraversal2(root *TreeNode) []int {
	var res []int
	return res
}
