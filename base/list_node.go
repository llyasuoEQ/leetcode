package base

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	temp := node
	res := ""
	for temp != nil {
		res += (fmt.Sprint(temp.Val) + "->")
		temp = temp.Next
	}
	res += "nil"
	return res
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}
