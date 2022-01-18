package search_insert_position

// searchInsert 搜索插入位置
// 因为有序时间复杂度要求是O(logn)所以利用二分查找法
// 时间复杂度： O(logn)
func searchInsert(nums []int, target int) int {
	length := len(nums)
	ret := length
	left, right := 0, length-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid - 1
			ret = mid
		} else {
			left = mid + 1
		}
	}
	return ret
}
