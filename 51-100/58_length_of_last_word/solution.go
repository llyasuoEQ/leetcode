package length_of_last_word

import "strings"

// lengthOfLastWord 最后一个单词长度
// 解题思路：从后往前遍历，首先删除末尾空格，然后从后往前遍历查找最后一个单词计算长度
func lengthOfLastWord(s string) int {
	// 去除字符串后面的空格
	s = strings.TrimRight(s, " ")

	// 从后往前遍历，查找到最后一个单词
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		length++
	}

	return length
}
