package wildcard_matching

// isMatch 通配符匹配
// 解题思路：动态规划解答，列出状态转移方程
// 假设dp[i][j]表示s的前i个字符与p的前j个字符是否匹配，i和j对应到s和p就是s[i-1]和p[j-1]
// 情况一：当s[i-1] == p[j-1] 或者 p[j] = "?" 时，那么dp[i][j] = dp[i-1][j-1]
// 情况二：当p[j-1] = "*" 时，表示p的第j个字符可以匹配s中的前的若干个字符
// ------ 当"*"没有匹配第i字符，那么dp[i][j] = dp[i-1][j]
// ------ 当"*"匹配上第i个字符，那么dp[i][j] = dp[i][j-1]
// ------ 这俩种情况满足一种即可，也就是要么匹配上要么匹配不上
func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)

	// 创建dp二维数组，dp[i][j]表示s的前i个字符与p的前j个字符是否匹配
	dp := make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true // 表示俩个空串匹配，表示true

	// 处理前置为"*"的情况，0的一行是s字符串无需处理，0的一列是p有"*"的情况，能匹配空串
	for i := 1; i <= pLen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	// 计算匹配逻辑
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[sLen][pLen]
}
