package valid_numbers

import (
	"strings"
	"unicode"
)

// isNumber 有效数字
// 解题思路：按照数学思路分情况讨论解决
// e和E后面必须跟着整数，且只能出现一次
// .后面必须跟着数字，且只能出现一次并且不能出现在e/E的后面
// +和-只能出现在最前面或者e/E后面
func isNumber(s string) bool {
	// 去除字符串两端的空格
	s = strings.TrimSpace(s)

	// 标记是否出现数字、小数点、指数符号（e/E）
	numSeen := false
	dotSeen := false
	eSeen := false
	numAfterE := true // e后边是否有整数，默认为true，当e/E存在的时候改为false

	// 遍历字符串的每一个字符，分情况套路
	for i, ch := range s {
		if unicode.IsDigit(ch) { // 是否为十进制数字
			numSeen = true
			numAfterE = true
		} else if ch == '.' {
			if dotSeen || eSeen { // 判断前边是否出现过.或出现过e/E
				return false
			}
			dotSeen = true
		} else if ch == 'e' || ch == 'E' {
			if eSeen || !numSeen { // 判断前边是否出现过e/E或前边没有出现数
				return false
			}
			eSeen = true
			numAfterE = false
		} else if ch == '-' || ch == '+' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' { // 判断是否在首位或者e/E后面
				return false
			}
		} else {
			return false
		}
	}

	return numSeen && numAfterE // 判断是否出现过数字，并且最后一个数字是否在指数符号之后
}
