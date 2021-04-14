package merge_ordered_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{
		Val:  0,
		Next: nil,
	}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	if l1 == nil {
		node.Next = l2
	}
	if l2 == nil {
		node.Next = l1
	}
	return node.Next
}
