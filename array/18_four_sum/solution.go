package four_sum

import "sort"

// 前后指针法 - O(N^2) 有些边界条件无法满足，放弃
// func fourSum(nums []int, target int) [][]int {
// 	var res [][]int
// 	length := len(nums)
// 	if length < 4 {
// 		return res
// 	}
// 	// 首先排序
// 	sort.Ints(nums)
// 	// 俩层for循环，分别使用前后指针法
// 	// 第一层指针 前指针动，后指针不动
// 	l1, r1 := 0, length-1
// 	for l1 < r1 {
// 		// 第二层指针，前后指针同时动
// 		targetTemp := target - nums[l1] - nums[r1]
// 		l2, r2 := l1+1, r1-1
// 		for l2 < r2 {
// 			targetTemp = targetTemp - nums[l2] - nums[r2]
// 			if targetTemp == 0 {
// 				item := []int{nums[l1], nums[l2], nums[r1], nums[r2]}
// 				res = append(res, item)
// 			}
// 			if targetTemp >= 0 {
// 				// 说明target < 四数之和，那么左指针就得增加，使得和尽量靠近target
// 				l2++
// 			} else {
// 				// 说明target > 四数之和，那么右指针就得减少，使得和尽量靠近target
// 				r2--
// 			}
// 		}
// 		if targetTemp >= 0 {
// 			// 说明target < 四数之和，那么左指针就得增加，使得和尽量靠近target，外层也是一样的道理
// 			// 因为最后的target值是最接近的值
// 			l1++
// 		} else {
// 			// 说明target > 四数之和，那么右指针就得减少，使得和尽量靠近target，外层也是一样的道理
// 			r1--
// 		}
// 	}
// 	return res
// }

// [1,0,-1,0,-2,2]
// [-2,-1,0,0,1,2]
// 前后指针法，三层循环去做
// 时间复杂度：O(N^3)
// 空间复杂度：O(1)
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	length := len(nums)
	sort.Ints(nums)
	if length < 4 {
		return res
	}
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, length-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					item := []int{nums[i], nums[j], nums[l], nums[r]}
					res = append(res, item)
					// 然后去判断l、r指针的指向
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				}
				if sum > target {
					r--
				}
				if sum < target {
					l++
				}
			}
		}
	}
	return res
}
