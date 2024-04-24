package base

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	temp := node
	res := ""
	loopMap := make(map[*ListNode]bool)
	for temp != nil {
		res += (fmt.Sprint(temp.Val) + "->")
		if _, ok := loopMap[temp]; ok { // 说明是环
			break
		}
		loopMap[temp] = true
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
