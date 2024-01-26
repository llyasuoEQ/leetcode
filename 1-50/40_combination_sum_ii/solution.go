package combination_sum_ii

import (
	"sort"
)

// combinationSum 组合总和
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var current []int
	sort.Ints(candidates)
	backtrack(candidates, &res, current, target, 0)
	return res
}

func backtrack(candidates []int, res *[][]int, current []int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, current...))
		return
	}
	for i := start; i < len(candidates); i++ {
		// 去除重复数字
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		current = append(current, candidates[i])
		backtrack(candidates, res, current, target-candidates[i], i+1)
		current = current[:len(current)-1]
	}
}
