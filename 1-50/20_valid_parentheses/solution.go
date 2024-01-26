package valid_parentheses

// isValid 利用栈来实现
func isValid(s string) bool {
	length := len(s)
	if length%2 > 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < length; i++ {
		if pairs[s[i]] > 0 { // 为右括号，查看是否有左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else { // 追加在左括号里
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
