package unique_paths

// uniquePaths 不通路径
// 解题思路：使用动态规划来解答，对于到达网格中的任意一点，要么是上边下来，要么是左边过来
// 那么到达finish的所有路径就是能到达上方和下方的所有不同路径的和
func uniquePaths(m int, n int) int {
	// 创建一个二维数组来保存路径数量
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一列和第一行的都位1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 动态规划计算路径数量
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
