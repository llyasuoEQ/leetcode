package recover_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：一般的解题思路
// 简单思路
// 找出需要交换的节点
// 然后交换
// 时间复杂度：O(3 * N) -> O(N)
// 空间复杂度：O(N)
func recoverTree1(root *TreeNode) {
	// 遍历树
	nums := inOrderTree(root)
	// 寻找需要交换的节点
	x, y := findTwoNode(nums)
	// 然后交换节点值
	recover(root, 2, x, y)
}

// 中序遍历
func inOrderTree(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, inOrderTree(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrderTree(root.Right)...)
	return
}

// 寻交换的节点值
func findTwoNode(nums []int) (x, y int) {
	x, y = -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return
}

// 交换节点值
func recover(root *TreeNode, count int, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Left, count, x, y)
	recover(root.Right, count, x, y)
}

// 方法二：
// 隐式的去中序遍历二叉树，然后找到需要交换的节点
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func recoverTree2(root *TreeNode) {
	stack := []*TreeNode{}
	temNode := root
	var x, y, pred *TreeNode
	for len(stack) > 0 || temNode != nil {
		for temNode != nil {
			stack = append(stack, temNode)
			temNode = temNode.Left
		}
		temNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && pred.Val > temNode.Val {
			y = temNode
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = temNode
		temNode = temNode.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

// 方法三：
// morris遍历
// 时间复杂度 O(N)
// 空间复杂度 O(1)
func recoverTree3(root *TreeNode) {
	// temNode := root
	// var x, y, pred *TreeNode
	// for temNode != nil {

	// }
}

// func recoverTree(root *TreeNode) {
// 	var x, y, pred, predecessor *TreeNode
// 	for root != nil {
// 		if root.Left != nil {
// 			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
// 			predecessor = root.Left
// 			for predecessor.Right != nil && predecessor.Right != root {
// 				predecessor = predecessor.Right
// 			}

// 			// 让 predecessor 的右指针指向 root，继续遍历左子树
// 			if predecessor.Right == nil {
// 				predecessor.Right = root
// 				root = root.Left
// 			} else { // 说明左子树已经访问完了，我们需要断开链接
// 				if pred != nil && root.Val < pred.Val {
// 					y = root
// 					if x == nil {
// 						x = pred
// 					}
// 				}
// 				pred = root
// 				predecessor.Right = nil
// 				root = root.Right
// 			}
// 		} else { // 如果没有左孩子，则直接访问右孩子
// 			if pred != nil && root.Val < pred.Val {
// 				y = root
// 				if x == nil {
// 					x = pred
// 				}
// 			}
// 			pred = root
// 			root = root.Right
// 		}
// 	}
// 	x.Val, y.Val = y.Val, x.Val
// }
