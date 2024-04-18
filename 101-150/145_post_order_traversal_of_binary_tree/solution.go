package post_order_traversal_of_binary_tree

import (
	"myProject/leetcode/base"
)

// 方法一：
// 递归实现
func postorderTraversal(root *base.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 方法二：
// 非递归实现，利用栈来实现
func postorderTraversal1(root *base.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := []*base.TreeNode{root}
	var pre *base.TreeNode

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if (node.Left == nil && node.Right == nil) || (pre != nil && (pre == node.Left || pre == node.Right)) { // 叶子节点
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			pre = node
		} else {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return res
}

// 方法三：
// 非递归实现，morris遍历
