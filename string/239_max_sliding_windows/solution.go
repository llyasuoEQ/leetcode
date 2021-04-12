package max_sliding_windows

import (
	"sort"
)

// 通常的解题思路
// 滑动窗口内的值求最大值，然后保存返回
func maxSlidingWindow1(nums []int, k int) []int {
	var res []int
	numsLen := len(nums)
	if k > numsLen || k < 1 {
		return res
	}
	left := 0
	right := k - 1
	for right < numsLen {
		right++
		windows := nums[left:right]
		sort.Ints(windows)
		res = append(res, windows[k-1])
		left++
	}
	return res
}

// 借助堆（优先队列）去实现
