package spiral_matrix

// spiralOrder 螺旋矩阵
// 解题思路：用螺旋遍历的思路来解答
func spiralOrder(matrix [][]int) []int {
	var result []int

	if len(matrix) == 0 {
		return result
	}

	// 定义四个变量来表示当前要遍历的范围
	var (
		rowBegin int = 0
		rowEnd   int = len(matrix) - 1
		colBegin int = 0
		colEnd   int = len(matrix[0]) - 1
	)

	// 按照螺旋遍历思路遍历
	for rowBegin <= rowEnd && colBegin <= colEnd {
		// 向右遍历，遍历完成后rowBegin加1
		for i := colBegin; i <= colEnd; i++ {
			result = append(result, matrix[rowBegin][i])
		}
		rowBegin++

		// 向下遍历，遍历完成后colEnd减1
		for i := rowBegin; i <= rowEnd; i++ {
			result = append(result, matrix[i][colEnd])
		}
		colEnd--

		// 满足rowBegin 小于等于 rowEnd时（否则会重复遍历），向左遍历，遍历完成后rowEnd减1，
		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				result = append(result, matrix[rowEnd][i])
			}
			rowEnd--
		}

		// 满足colBegin 小于等于 colEnd时（否则会重复遍历），向上遍历，遍历完成后colBegin加1
		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				result = append(result, matrix[i][colBegin])
			}
			colBegin++
		}
	}
	return result
}
