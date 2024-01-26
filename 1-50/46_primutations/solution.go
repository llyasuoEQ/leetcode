package primutations

// permute 全排列
// 解题思路：利用回溯算法来实现
func permute(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, []int{}, &result)
	return result
}

// backtrack 回溯算法
func backtrack(nums []int, current []int, result *[][]int) {
	if len(nums) == len(current) {
		// 当排列长度等于原数组长度的时候，将排列结果添加至结果中
		*result = append(*result, append([]int{}, current...))
		return
	}
	for _, num := range nums {
		if contains(current, num) {
			continue
		}
		// 将当前数字添加至当前排列current中
		current = append(current, num)
		// 继续递归生成下一个数字排列
		backtrack(nums, current, result)
		// 回溯，将当前数组从排列中移除
		current = current[:len(current)-1]
	}

}

// contains 判断数组中是否包含某个元素
func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
