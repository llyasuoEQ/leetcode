package removen_th_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

// removeNthFromEnd 遍历俩遍实现
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 确定有多少节点
	var (
		length int
		temp   = head
	)
	for temp != nil {
		if temp != nil {
			length++
		}
		temp = temp.Next
	}
	delPoint := length - n // 计算删除的前置节点
	if delPoint < 0 {
		return nil
	}
	if delPoint == 0 { // 删除的是头结点
		head = head.Next
	} else { // 删除非头节点
		temp = head
		for i := 1; i <= delPoint; i++ {
			if i == delPoint { // 删除该节点
				if temp.Next.Next == nil {
					temp.Next = nil
				} else {
					temp.Next = temp.Next.Next
				}
				break
			}
			temp = temp.Next
		}
	}
	return head
}

// TODO 进阶，一遍遍历完成删除
