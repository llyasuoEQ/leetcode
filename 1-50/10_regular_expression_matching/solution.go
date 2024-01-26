package regular_expression_matching

// isMatch 动态规划来解
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' {
				if i > 0 {
					f[i][j] = f[i-1][j-1]
				} else {
					f[i][j] = false
				}
			} else if p[j-1] == '*' {
				if (i > 0) && (s[i-1] == p[j-2] || p[j-2] == '.') {
					f[i][j] = f[i-1][j] || f[i][j-2]
				} else {
					f[i][j] = f[i][j-2]
				}
			} else {
				if i > 0 && (s[i-1] == p[j-1]) {
					f[i][j] = f[i-1][j-1]
				} else {
					f[i][j] = false
				}
			}
		}
	}
	return f[m][n]
}
