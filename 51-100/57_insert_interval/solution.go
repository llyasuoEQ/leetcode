package insert_interval

// insert 插入区间
// 解题思路：使用线性扫描（也就是依次扫描数组，按照规则处理）的方法来解答
func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		result   [][]int // 输出结果
		inserted bool    // 标记是否插入成功
	)

	for _, interval := range intervals {
		// 当前区域在待插区域的右侧，且没有交集
		if interval[1] < newInterval[0] {
			result = append(result, interval)
		} else if interval[0] > newInterval[1] { // 当前区域在待插区间的左侧且无交集
			if !inserted {
				result = append(result, newInterval)
				inserted = true
			}
			result = append(result, interval)
		} else { // 当前区间与待插入区间有交集，合并区间
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	if !inserted {
		result = append(result, newInterval)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
