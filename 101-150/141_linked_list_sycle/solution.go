package linked_list_sycle

import "myProject/leetcode/base"

// hasCycle 环形链表
// 解题思路：一种是使用map，但是有O(N)的时间复杂度，使用快慢指针的方法，当快指针追上慢指针时说明就是环
func hasCycle(head *base.ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for slow != nil && slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
