package verify_binary_search_tree_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯法
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var allTreeNode []*TreeNode
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currNode := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				allTreeNode = append(allTreeNode, currNode)
			}
		}
	}
	return allTreeNode
}
