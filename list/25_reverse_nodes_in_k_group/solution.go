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

// reverseKGroup 翻转链表中K个为一组的节点
// 1. 先统计链表中需要翻转的子链表
// 2. 翻转子链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	// 1. 先统计节点的个数
	var (
		currentNode = head
		length      int
	)
	for currentNode != nil {
		length++
		currentNode = currentNode.Next
	}
	// 2. 统计需要翻转K个为一组的节点的次数
	reverseCount := length / k

	// 3. 分别翻转k个为一组的节点
	tempNode := &ListNode{Val: 0, Next: head}
	currentNode = tempNode
	for i := 0; i < reverseCount; i++ {
		var (
			startNode = currentNode.Next
			endNode   = startNode
		)
		for j := 1; j < k; j++ {
			endNode = endNode.Next
		}
		nextCurrentNode := endNode.Next
		endNode.Next = nil
		currentNode.Next = reverseList(startNode)
		currentNode = startNode
		startNode.Next = nextCurrentNode
	}
	return tempNode.Next
}

// reverseList 翻转链表
func reverseList(head *ListNode) *ListNode {
	var (
		currentNode = head
		preNode     *ListNode
	)
	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = preNode
		preNode = currentNode
		currentNode = next
	}
	return preNode
}
