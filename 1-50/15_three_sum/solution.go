package three_sum

import (
	"sort"
)

// 方法一：
// 暴力求解-显然这个时间复杂度太大，实在不会想出来的才会这样去做，当数据量实在大的时候，肯定A不过去，超时限制
// 时间复杂度：O(N^3)
// 空间复杂度：O(1)
func threeSum1(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}

	for i := 0; i < length-2; i++ {
		ti := 0 - nums[i]
		for j := i + 1; j < length-1; j++ {
			tj := ti - nums[j]
			for n := j + 1; n < length; n++ {
				if (tj - nums[n]) == 0 {
					temp := []int{nums[i], nums[j], nums[n]}
					// 去重
					if !isIndexOfArray(temp, res) {
						res = append(res, temp)
					}
				}
			}
		}
	}
	return res
}

// isIndexOfRes...
func isIndexOfArray(cur []int, arr [][]int) bool {
	for _, v := range arr {
		indexMap := map[int]int{}
		for _, vv := range v {
			_, ok := indexMap[vv]
			if ok {
				indexMap[vv]++
			} else {
				indexMap[vv] = 1
			}
		}
		for _, cv := range cur {
			value, ok := indexMap[cv]
			if !ok {
				break
			}
			if value > 1 {
				indexMap[cv]--
			} else {
				delete(indexMap, cv)
			}
		}
		if len(indexMap) == 0 {
			return true
		}
	}
	return false
}

// 方法二：
// 排序+双指针法
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func threeSum2(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		// 因为排序好的，所以当前值和前一个值不能相同，这样可以去除重复一部分
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// r 对应指针的最右边
		r := length - 1
		target := -1 * nums[first]
		for l := first + 1; l < length; l++ {
			if l > first+1 && nums[l] == nums[l-1] {
				continue
			}
			for l < r && nums[l]+nums[r] > target {
				r--
			}
			if l == r {
				break
			}
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[first], nums[l], nums[r]})
			}
		}
	}

	return res
}
