package convert_sorted_list_to_binary_search_tree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedListToBST1 有序链表转换为高度平衡二叉搜索树
// 分治方法
// 时间复杂度：O(nlogn)
// 空间复杂度；O(logn)
func sortedListToBST1(head *ListNode) *TreeNode {
	return buildBST1(head, nil)
}

// getListMid 寻找列表中位节点
// 快慢指针法
func getListMid(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// buildBST 构建平衡二叉搜索树
func buildBST1(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getListMid(left, right)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = buildBST1(left, mid)
	root.Right = buildBST1(mid.Next, right)
	return root
}

var globalHead *ListNode

// sortedListToBST2 分治+中序遍历
// 方法一的性能瓶颈是每次都需要把链表遍历一遍，因此可以把分治和中序遍历结合起来减少时间复杂度
func sortedListToBST2(head *ListNode) *TreeNode {
	globalHead = head
	length := getListLength(head)
	return buildBST2(0, length-1)
}

// getListLength 获取链表长度
func getListLength(head *ListNode) int {
	var ret int
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildBST2(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &TreeNode{}
	root.Left = buildBST2(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildBST2(mid+1, right)
	return root
}
