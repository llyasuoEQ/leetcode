package find_first_and_last_position_of_element_insorted_array

// 方法一：利用二分查找
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			// 找到了
			tL, tR := mid, mid
			// 找出目标值的左右下标
			for tL > -1 && nums[tL] == target {
				tL--
			}
			for tR < len(nums) && nums[tR] == target {
				tR++
			}
			return []int{tL + 1, tR - 1}
		}
	}
	return []int{-1, -1}
}
