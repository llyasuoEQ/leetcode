package minimum_paths_sum

// minPathSum 最小路径和
// 解题思路：使用动态规划
// 设dp为大小m*n矩阵，其中dp[i][j]的值代表走到位置(i,j)的最小路径和
// 题目中要求只能向下或者向左走，那么只考虑左边界和右边界
// 那么状态转移方程为：dp[i][j] = min(dp[i-1][j],dp[i][j-1])，边界时特殊考虑
func minPathSum(grid [][]int) int {
	m := len(grid) // 长
	n := len(grid[0])

	// 创建一个二维数组来保存最小路径和
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一列的最小路径和，累加grid第一列
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 初始化第一行最小路径和，累加grid第一行
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 动态规划计算最小路径和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
