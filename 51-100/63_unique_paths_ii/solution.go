package unique_paths_ii

// uniquePathsWithObstacles 不同路径ii
// 解题思路：利用动态规划，和61不同路径的解法是一样的，但是这里考虑障碍物的情况
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)    // 高
	n := len(obstacleGrid[0]) // 宽

	// 创建一个二维数组来保存路径数量
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一列和第一行的路径数量，当这一列或者这一行有障碍物就结束，因为只能向下或者向右所以有障碍物就过不去了
	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}

	// 动态规划计算路径数量
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}
