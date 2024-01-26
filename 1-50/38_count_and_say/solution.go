package count_and_say

import "strconv"

// countAndSay 外观数列，利用递归来实现
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	preSay := countAndSay(n - 1)
	var (
		curSay string
		count  int = 1
	)
	// 遍历上一个数列，实现下一个数列
	for i := 0; i < len(preSay); i++ {
		if i+1 < len(preSay) && preSay[i] == preSay[i+1] { // 当前值与后置值相等，计数加一
			count++
		} else { // 当前值与后置值不相等，count重置同时组合字符串
			curSay += strconv.Itoa(count) + string(preSay[i])
			count = 1
		}
	}

	return curSay
}

// countAndSay2 外观数列，利用非递归的方法实现
func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}

	var result = "1"
	for i := 2; i < n+1; i++ {
		var (
			resultTemp string
			count      = 1
		)
		for j := 0; j < len(result); j++ {
			if j+1 < len(result) && result[j] == result[j+1] {
				count++
			} else {
				resultTemp += strconv.Itoa(count) + string(result[j])
				count = 1
			}
		}
		result = resultTemp
	}
	return result
}
