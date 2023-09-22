package merge_k_sorted_lists

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
	return res
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

// mergeTwoList 合并俩个升序列表
func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	for list1 != nil && list2 != nil {
		fmt.Println("====", list1, list2)
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

// mergeKLists2 使用分治思路解决
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	return mergeBacktrack(lists, 0, len(lists)-1)
}

func mergeBacktrack(lists []*ListNode, left, right int) *ListNode {
	if right-left < 1 {
		return lists[left]
	}
	mid := (left + right) / 2
	list1 := mergeBacktrack(lists, 0, mid)
	list2 := mergeBacktrack(lists, mid+1, right)
	return mergeTwoList(list1, list2)
}
