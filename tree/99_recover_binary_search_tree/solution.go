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
// func recoverTree2(root *TreeNode) {
// 	stack := []*TreeNode{}
// }
