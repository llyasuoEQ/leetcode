package reverse_nodes_in_k_group

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var res string
	for l != nil {
		res = res + fmt.Sprint(l.Val) + "->"
		l = l.Next
	}
	res += "nil"
	return res
}

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	return nil
}
