package preorder_traversal_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 递归实现
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// TODO 迭代实现
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	return res
}

// 方法二：
// 非递归实现

// 方法三：
// 非递归实现，morris遍历
