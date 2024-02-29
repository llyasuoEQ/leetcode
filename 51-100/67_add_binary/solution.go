package add_binary

import "strconv"

// addBinary 二进制求和
// 解题思路：从后往前遍历俩个字符串，依次相加
func addBinary(a string, b string) string {
	var (
		result string
		carry  int                      // 进位
		i, j   = len(a) - 1, len(b) - 1 // 指向a和b的末尾
	)

	// 遍历a和b从字符串最后一位开始依次相加
	for i >= 0 || j >= 0 || carry > 0 {
		// 获取当前位置的值，如果字符串已经遍历完成，则取0
		digitA, digitB := 0, 0
		if i >= 0 {
			digitA = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			digitB = int(b[j] - '0')
			j--
		}

		// 计算当前位的和以及进位
		sum := digitA + digitB + carry
		currDight := sum % 2
		carry = sum / 2

		// 将当前位的和插入结果字符串的最前面
		result = strconv.Itoa(currDight) + result
	}

	return result
}
