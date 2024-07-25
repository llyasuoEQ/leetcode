package linked_list_cycle_ii

import (
	"myProject/leetcode/base"
)

// detectCycle 环形链表II
// 解题思路：使用快慢指针来执行
func detectCycle(head *base.ListNode) *base.ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 证明有环
			pre := head
			for pre != slow {
				pre = pre.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
