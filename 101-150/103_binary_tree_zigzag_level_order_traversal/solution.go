package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1.层序遍历的思路
// 2.遍历层级
// - 偶数层 无需变化
// - 奇数层 左右值对调
// 时间复杂度：O(N)，其中 NN 为二叉树的节点数。每个节点会且仅会被遍历一次。
// 空间复杂度：O(N)。我们需要维护存储节点的队列和存储节点值的双端队列，空间复杂度为 O(N)O(N)
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	currQ := []*TreeNode{root}
	// i 代表层级
	for i := 0; len(currQ) > 0; i++ {
		res = append(res, []int{})
		nextQ := []*TreeNode{}
		for j := 0; j < len(currQ); j++ {
			node := currQ[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}
		currQ = nextQ
		if i%2 == 1 {
			l, r := 0, len(res[i])-1
			for l < r {
				temp := res[i][l]
				res[i][l] = res[i][r]
				res[i][r] = temp
				l++
				r--
			}
		}
	}
	return res
}
