package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 递归实现
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	pLen := len(postorder)
	root := &TreeNode{
		Val:   postorder[pLen-1],
		Left:  nil,
		Right: nil,
	}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == postorder[pLen-1] {
			break
		}
	}
	root.Left = buildTree1(inorder[:i], postorder[:i])
	root.Right = buildTree1(inorder[i+1:], postorder[i:pLen-1])
	return root
}

// 方法二：
// 迭代实现-借助栈实现
// 时间复杂度：
// 空间复杂度：
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}
	stack := []*TreeNode{root}
	index := len(inorder) - 1
	for i := len(postorder) - 2; i >= 0; i-- {
		node := stack[len(stack)-1]
		if node.Val != inorder[index] {
			// 左子树节点
			node.Right = &TreeNode{
				Val:   postorder[i],
				Left:  nil,
				Right: nil,
			}
			stack = append(stack, node.Right)
		} else {
			// 判断是谁的右子树
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index--
			}
			// 右子树节点
			node.Left = &TreeNode{
				Val:   postorder[i],
				Left:  nil,
				Right: nil,
			}
			stack = append(stack, node.Left)
		}
	}

	return root
}
