package preorder_traversal_binary_tree

import "myProject/leetcode/base"

// 方法一：
// 递归实现
func preorderTraversal(root *base.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// TODO 迭代实现，利用栈先进先出的思路解决
func preorderTraversal2(root *base.TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*base.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		// 先入先出，先放右子树，然后放左子树
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

// 方法三：
// 非递归实现，morris遍历
