package reorder_list

import (
	"myProject/leetcode/base"
)

// reorderList 重新排序链表 1 2 3 4 -> 1 4 2 3
// 解题思路：拆分链表、反转链表，合并链表
func reorderList(head *base.ListNode) {
	if head == nil {
		return
	}

	// 寻找中间节点
	mid := middleNode(head)

	// 拆分
	l1 := head
	l2 := mid.Next
	mid.Next = nil

	// 反转l2
	l2 = reverseList(l2)

	mergeList(l1, l2)
}

// middleNode 寻找中间节点，使用快慢制作
func middleNode(head *base.ListNode) *base.ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// reverseList 反转链表
func reverseList(head *base.ListNode) *base.ListNode {
	var pre, cur *base.ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTemp
	}
	return pre
}

// mergeList 合并链表，比如1，3，5和2，4，6 => 1,2,3,4,5,6
func mergeList(l1, l2 *base.ListNode) {
	var l1Temp, l2Temp *base.ListNode
	for l1 != nil && l2 != nil {
		l1Temp = l1.Next
		l2Temp = l2.Next

		l1.Next = l2
		l1 = l1Temp

		l2.Next = l1
		l2 = l2Temp
	}
}

// reorderList2 解决方法二，利用数组访问下表的方式
func reordrList2(head *base.ListNode) {
	if head == nil {
		return
	}

	// 存放在数组中
	var list []*base.ListNode
	temp := head
	for temp != nil {
		list = append(list, temp)
		temp = temp.Next
	}

	// 前后指针，依次组合
	var left, right = 0, len(list) - 1
	for left < right {
		list[left].Next = list[right]
		left++
		if left == right {
			break
		}

		list[right].Next = list[left]
		right--
	}
	list[left].Next = nil
}
