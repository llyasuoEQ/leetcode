package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}
	if list1 == nil {
		node.Next = list2
	}
	if list2 == nil {
		node.Next = list1
	}

	return head.Next
}
