package combination_sum

// combinationSum 组合总和
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var current []int
	backtrack(&res, candidates, target, current, 0)
	return res
}

// backtrack 回溯逻辑
func backtrack(res *[][]int, candidates []int, target int, current []int, start int) {
	if target < 0 {
		return
	}
	if target == 0 { // 表示当前满足规则，直接追加到结果中
		*res = append(*res, append([]int{}, current...))
	}
	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(res, candidates, target-candidates[i], current, i)
		// 上一轮递归处理完成了，那么就删除掉current元素
		current = current[:len(current)-1]
	}

}
