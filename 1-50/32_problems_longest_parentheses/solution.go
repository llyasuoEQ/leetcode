package problems_longest_parentheses

// longestValidParentheses 最长有效括号
// 使用栈来后进先出解决该问题，栈里来记录左括号，当有右括号时弹出栈顶元素
func longestValidParentheses(s string) int {
	// 创建栈
	stack := []int{-1}
	maxLen := 0

	for i, char := range s {
		if char == '(' { // 左括号入栈
			stack = append(stack, i)
		} else { // 右括号出栈
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 { // 如果栈为空，放入当前位置记录
				stack = append(stack, i)
			} else { // 计算最大有效长度
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
