package in_order_traversal_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左 -> 根 -> 右
// 方法一：
// 递归的方式实现
// 时间复杂度：O(N)
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)
	return res
}

// 方法二：
// 遍历的方式
// 利用栈的思想：后进先出
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	temNode := root
	stack := []*TreeNode{}
	for len(stack) > 0 || temNode != nil {
		for temNode != nil {
			stack = append(stack, temNode)
			temNode = temNode.Left
		}
		temNode = stack[len(stack)-1]
		res = append(res, temNode.Val)
		stack = stack[:len(stack)-1]
		temNode = temNode.Right
	}
	return res
}
