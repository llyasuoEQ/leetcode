package slimbing_stairs

// climbStairs 爬楼梯
// 解题思路：使用动态规划来解决
// dp[i]表示爬到第i阶梯的不同方法，那么dp[i] = dp[i-1] + dp[i-2]
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
