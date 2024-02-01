package merge_intervals

import "sort"

// merge 合并区间
// 解题思路：使用排序和合并区间的方法来解答
func merge(intervals [][]int) [][]int {
	// 对区间起始进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		curr := intervals[i]

		// 当前区间的起始位置在上一个合并区间范围内
		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
