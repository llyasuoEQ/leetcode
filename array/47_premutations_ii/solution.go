package premutations_ii

import (
	"sort"
)

// permuteUnique 全排列2
// 解题思路：利用回溯算法解答
func permuteUnique(nums []int) [][]int {
	// 没有说是排序好的数组，那么先排序
	sort.Ints(nums)
	result := [][]int{}
	// 利用下表来标记nums中的元素是否被访问过
	visited := make([]bool, len(nums))
	backtrack(nums, []int{}, visited, &result)
	return result
}

// backtrack 回溯方法
func backtrack(nums []int, current []int, visited []bool, result *[][]int) {
	if len(nums) == len(current) {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for i := 0; i < len(nums); i++ {
		// 当前元素被使用过或者与前一个元素相同并且未被使用，则跳过当前元素
		if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
			continue
		}
		// 标记当前元素访问过了
		visited[i] = true
		// 将当前元素添加到当前排列中
		current = append(current, nums[i])
		// 继续递归生产下一个数字排列
		backtrack(nums, current, visited, result)

		// 回溯，将当前元素只为未访问，同时将当前元素从排列中移除
		visited[i] = false
		current = current[:len(current)-1]
	}
}
