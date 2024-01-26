package divide_two_integers

import "math"

// divide 先计算结果的正负，然后用被除数依次减去除数
// 这个解法当被除数较大除数较小的时候会超出时间限制
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// dividend，divisor 同号为正，不同号为负
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}

	// dividend和divisor分别取绝对值，依次对dividend减去divisor，
	// 直到dividend < divisor，返回截取的次数 * sign
	var (
		result      = 0
		absDividend = abs(dividend)
		absDivisor  = abs(divisor)
	)

	for absDividend >= absDivisor {
		absDividend -= absDivisor
		result++
	}
	if sign < 0 {
		result = -result
	}
	return result
}

// divide2 先计算结果的正负，然后用被除数依次减去除数
func divide2(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// dividend，divisor 同号为正，不同号为负
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}

	// dividend和divisor分别取绝对值，依次对dividend减去divisor，
	// 直到dividend < divisor，返回截取的次数 * sign
	var (
		result      = 0
		absDividend = abs(dividend)
		absDivisor  = abs(divisor)
	)

	// 这里用二分查找法，不然会超出时间限制
	for absDividend >= absDivisor {
		temp, multiple := absDivisor, 1
		for absDividend >= (temp << 1) {
			temp = temp << 1
			multiple = multiple << 1
		}
		absDividend -= temp
		result += multiple
	}
	if sign < 0 {
		result = -result
	}
	return result
}

// abs 取绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
