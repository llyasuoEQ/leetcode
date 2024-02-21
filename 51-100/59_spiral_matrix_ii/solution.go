package spiral_matrix_ii

// generateMatrix 螺旋矩阵II
// 解题思路：使用代码模拟螺旋便利的方法
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 存放的旋转数字
	num := 1
	top, bottom, left, right := 0, n-1, 0, n-1

	for num <= n*n {
		// 从左往右
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// 从上往下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// 从右往左
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// 从下往上
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
