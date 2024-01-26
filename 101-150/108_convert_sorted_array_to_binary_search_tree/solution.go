package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现-中序遍历取中间节点为根节点，具体取中间左边还是右边节点都可以
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func sortedArrayToBST(nums []int) *TreeNode {
	nLen := len(nums)
	if nLen == 0 {
		return nil
	}
	median := nLen / 2
	root := &TreeNode{
		Val:   nums[median],
		Left:  nil,
		Right: nil,
	}
	root.Left = sortedArrayToBST(nums[:median])
	root.Right = sortedArrayToBST(nums[median+1:])
	return root
}
