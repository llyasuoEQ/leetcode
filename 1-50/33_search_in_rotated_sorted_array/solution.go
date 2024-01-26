package search_in_rotated_sorted_array

// [4,5,6,7,0,1,2]
// 方法一: 利用二分查找的方式 - 这种方案问题考虑的边界条件较多，容易写错
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	res := -1
	length := len(nums)
	if length < 1 {
		return res
	}
	l, r := 0, len(nums)-1
	for l <= r {
		div := nums[l] // 记录下最左侧的值，旋转的分界值
		medain := (l + r) / 2
		// 找到了，直接return
		if target == nums[medain] {
			res = medain
			break
		}
		// 首先判断中间数是否越过分界线
		if nums[medain] < div {
			// 中间数越过分界线场景
			if target < div && target > nums[medain] {
				l = medain + 1
			} else {
				r = medain - 1
			}
		} else {
			// 中间数未越过分界线场景
			if target < nums[medain] && target >= div {
				r = medain - 1
			} else {
				l = medain + 1
			}
		}
	}
	return res
}
