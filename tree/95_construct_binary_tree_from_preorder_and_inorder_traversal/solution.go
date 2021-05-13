package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 递归实现
// 实现左右字数
// 前序遍历可以定位到根节点
// 中序遍历可以定位当前节点左右子树的数目
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	node.Left = buildTree1(preorder[1:i+1], inorder[:i])
	node.Right = buildTree1(preorder[i+1:], inorder[i+1:])
	return node
}

// 方法二：
// 迭代方式实现
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	inorderIndex := 0
	stack := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{
				Val:   preorderVal,
				Left:  nil,
				Right: nil,
			}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{
				Val:   preorderVal,
				Left:  nil,
				Right: nil,
			}
			stack = append(stack, node.Right)
		}
	}
	return root
}
