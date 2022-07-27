package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// 计算俩数之和
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	jw := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		tmp := v1 + v2 + jw
		v := tmp % 10
		jw = tmp / 10
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next
	}
	if jw > 0 {
		head.Next = &ListNode{
			Val: jw,
		}
	}

	return tmp.Next
}
