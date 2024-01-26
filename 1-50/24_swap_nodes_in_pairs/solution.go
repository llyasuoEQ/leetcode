package swap_nodes_in_pairs

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

// swapPairs 遍历一遍交换
func swapPairs(head *ListNode) *ListNode {
	var (
		node      = head // 临时节点
		i         = 0    // 记录位置，当每次 i%2 == 0 是需要交换节点
		node1     *ListNode
		indexNode *ListNode
	)
	for node != nil {
		i++
		if i%2 == 1 {
			node1 = node
		}
		if i%2 == 0 { // 达到交换的条件
			if i == 2 { // 头结点交换
				head = node
			} else {
				indexNode.Next = node
			}
			temp := node.Next
			node.Next = node1
			node1.Next = temp
			indexNode = node1
			node = node1
		}
		node = node.Next
	}
	return head
}

// swapPairs1 递归的方式解决
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs1(next.Next)
	next.Next = head
	return next
}
