package binary_tree_level_order_traversal_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 遍历出二叉树的层序遍历结果 -> 反转结果数组
// 时间复杂度：O（h * N）,h为层数
// 空间复杂度：O（N）
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue1 := []*TreeNode{root}
	level := 0
	for len(queue1) > 0 {
		queue2 := []*TreeNode{}
		res = append(res, []int{})
		for i := 0; i < len(queue1); i++ {
			node := queue1[i]
			res[level] = append(res[level], node.Val)
			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}
		queue1 = queue2
		level++
	}
	res = swap(res)
	return res
}

func swap(arr [][]int) [][]int {
	b, a := 0, len(arr)-1
	for b < a {
		temp := arr[b]
		arr[b] = arr[a]
		arr[a] = temp
		b++
		a--
	}
	return arr
}
