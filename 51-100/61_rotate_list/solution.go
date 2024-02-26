package rotate_list

import "fmt"

// Definition for singly-linked list.
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

// rotateRight 旋转链表
// 解题思路：使用链表操作来解决，首先将链表做成环，然后从最后的节点开始遍历找到要反转的节点，然后断开节点
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算链表长度
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// 将尾节点连接到头节点
	tail.Next = head

	// 计算实际要旋转的补助
	k %= length

	// 找到新的头节点的前一个节点
	newHeadPrev := tail
	for i := 0; i < length-k; i++ {
		newHeadPrev = newHeadPrev.Next
	}

	// 断开循环链表，将新的头节点与前一个节点连接
	newHead := newHeadPrev.Next
	newHeadPrev.Next = nil

	return newHead
}
