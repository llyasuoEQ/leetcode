package remove_element

// [1, 2, 2, 3, 4]
// 前后指针法
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == val {
			nums[l] = nums[r-1]
			r--
		} else {
			l++
		}
	}
	return l
}
