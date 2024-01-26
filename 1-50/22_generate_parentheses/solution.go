package generate_parentheses

// generateParenthesis 括号生成
// 从小变多
// 当n == 1 "()"
// 当n > 2 是，每次n+1，就是在原有的基础左边增加()、左边增加()、括号包含(XX)，然后再利用map去重
// 这种方法经测验是错误，有部分情况不能满足，比如n=4时，(())(())
// func generateParenthesisXX(n int) []string {
// 	if n < 1 {
// 		return nil
// 	}
// 	parenthesis := []string{"()"}
// 	if n == 1 {
// 		return parenthesis
// 	}
// 	duplicationCheck := make(map[string]bool)
// 	for i := 2; i <= n; i++ {
// 		var temp []string
// 		for _, item := range parenthesis {
// 			// 原有基础上用()包含
// 			include := "(" + item + ")"
// 			if !duplicationCheck[include] {
// 				temp = append(temp, include)
// 			}
// 			duplicationCheck[include] = true

// 			// 原有基础左边插入()
// 			leftInset := "()" + item
// 			if !duplicationCheck[leftInset] {
// 				temp = append(temp, leftInset)
// 			}
// 			duplicationCheck[leftInset] = true

// 			// 原有基础右边插入()
// 			rightInset := item + "()"
// 			if !duplicationCheck[rightInset] {
// 				temp = append(temp, rightInset)
// 			}
// 			duplicationCheck[rightInset] = true
// 		}
// 		parenthesis = temp
// 	}
// 	return parenthesis
// }

// generateParenthesis 利用递归来做，看成子问题，左边匹配右边
func generateParenthesis(n int) []string {
	var result []string
	generate("", 0, 0, n, &result)
	return result
}

func generate(current string, leftCount int, rightCount int, n int, result *[]string) {
	// 结束条件
	if len(current) == 2*n {
		*result = append(*result, current)
		return
	}

	// 左括号
	if leftCount < n {
		generate(current+"(", leftCount+1, rightCount, n, result)
	}

	// 右括号
	if rightCount < leftCount {
		generate(current+")", leftCount, rightCount+1, n, result)
	}
}
