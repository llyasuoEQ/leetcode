package string_to_integer_atoi

import "math"

// myAtoi ...
func myAtoi(s string) int {
	if len(s) < 1 {
		return 0
	}
	var (
		numType = 1 // 表示正数还是负数，1代表正数，-1代表负数
		price   = 0 // 输出的值
	)
	// 过滤空格
	for i, v := range s {
		if v != ' ' {
			s = s[i:]
			break
		}
	}
	// 检查正负号，并累计值
	for i, v := range s {
		if i == 0 && v == '-' {
			numType = -1
		} else if i == 0 && v == '+' {
			numType = 1
		} else if (v-'0') >= 0 && (v-'0') < 10 {
			price = price*10 + int(v-'0')
		} else {
			break
		}

		if numType*price < math.MinInt32 {
			return math.MinInt32
		}
		if numType*price > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return price * numType
}
