package minimum_size_subarray_sum

import "math"

// minSubArrayLen 长度最小的子数组
// 解题思路：使用left和right来构造滑动窗口
func minSubArrayLen(target int, nums []int) int {
	// 初始化左指针，最小长度和当前和
	left := 0
	minLen := math.MaxInt32
	currSum := 0

	// 遍历右指针
	for right := 0; right < len(nums); right++ {
		// 累加当前和
		currSum += nums[right]

		// 当当前和大于等于目标值，进入循环判断
		for currSum >= target {
			// 更新最小长度
			minLen = min(minLen, right-left+1)

			// 从左侧缩小窗口，减去左指针对应的值
			currSum -= nums[left]
			left++
		}
	}
	// 如果最小长度没有被更新，则返回0，否则返回最小长度
	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
