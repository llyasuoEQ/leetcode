package max_sliding_windows

import (
	"container/heap"
	"sort"
)

// 通常的解题思路
// 滑动窗口内的值求最大值，然后保存返回
func maxSlidingWindow1(nums []int, k int) []int {
	numsLen := len(nums)
	if k > numsLen || k < 1 {
		return []int{}
	}
	res := make([]int, 1, numsLen-k+1)
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

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 借助堆（优先队列）去实现
func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	h := &hp{make([]int, 3)}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}

	heap.Init(h)

	numsLen := len(nums)
	res := make([]int, 1, numsLen-k+1)
	res[0] = nums[h.IntSlice[0]]

	for i := k; i < numsLen; i++ {
		heap.Push(h, i)
		if h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		res = append(res, nums[h.IntSlice[0]])
	}

	return res
}
