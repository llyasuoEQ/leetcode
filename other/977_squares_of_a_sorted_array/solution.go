package squares_of_a_sorted_array

// sortedSquares 有序数组的平方
// 解题思路：使用双指针来解答
func sortedSquares(nums []int) []int {
	length := len(nums)
	left := 0
	right := length - 1
	index := length - 1

	result := make([]int, length)
	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			result[index] = nums[left] * nums[left]
			left++
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
