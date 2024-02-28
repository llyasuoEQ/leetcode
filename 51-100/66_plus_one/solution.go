package plus_one

// plusOne 加一
// 解题思路：分情况讨论
// 1. 不进位直接末尾+1
// 2. 进位
// - 进位处变为0，前一位+1
// - 位数不够在数组最前面添加1
func plusOne(digits []int) []int {
	length := len(digits)
	// 从数组的最后一位开始向前遍历
	for i := length - 1; i >= 0; i-- {
		// 如果当前位小于9，直接加一返回结果
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 如果当前位是9，将其变为0，然后继续向前遍历
		digits[i] = 0
	}
	// 如果整个数组都是9，那么需要在数组最前面插入一个1
	digits = append([]int{1}, digits...)
	return digits
}
