package merge_k_sorted_lists

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

// mergeKLists 暴力求解直接看成俩俩合并，当然这种算法时间复杂度较高
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	for _, list := range lists {
		head.Next = mergeTwoList(head.Next, list)
	}
	return head.Next
}

// mergeKLists2 使用分治思路解决
func mergeKLists2(lists []*ListNode) *ListNode {
	return nil
}

// mergeTwoList 合并俩个升序列表
func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
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
