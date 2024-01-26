package three_sum_closest

import (
	"math"
	"sort"
)

// 前后指针法
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func threeSumClosest(nums []int, target int) int {
	var res int
	diff := math.MaxInt32 // 先去最大值
	length := len(nums)
	if length < 3 {
		return res
	}
	sort.Ints(nums)
	for first := 0; first < length; first++ {
		// 前后指针
		r := length - 1
		l := first + 1
		for l < r {
			resTem := nums[first] + nums[l] + nums[r]
			// 确定是指针的移动
			if resTem > target {
				r--
			} else {
				l++
			}
			diffTemp := int(math.Abs(float64(target - resTem)))
			// 确定差值和结果
			if diffTemp < diff {
				res = resTem
				diff = diffTemp
			}
		}
	}
	return res
}

// 0
// 0 - 1 = -1
// 0 - (-1) = 1

// 正数
// 1
// 1 - 2 = -1
// 1 - 0 = 1
// 1 - (-2) = 3

// 负数
// -2 - 2 = -4
// -2 - (-4) = 2

// [-2,-1,0,1,2]
// 2 - [-2,-1,2] = 3
