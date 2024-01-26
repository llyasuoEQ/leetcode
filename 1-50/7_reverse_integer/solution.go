package reverse_integer

import "math"

// reverse 整数翻转，暴力求解
func reverse(x int) int {
	res := 0
	// 首先寻找从多少开始除
	temp := x
	divisor := 1
	for temp > 9 || temp < -9 {
		temp = x / divisor
		divisor *= 10
	}
	if divisor == 1 {
		return x
	}
	num := 0
	rDivisor := 1
	for divisor > 1 {
		divisor = divisor / 10
		num = x / divisor
		x = x % divisor
		res += rDivisor * num
		rDivisor *= 10
	}
	if res > math.MaxInt32/10 || res > math.MinInt32/10 {
		return 0
	}
	return res
}

// reverse1 优化解法
func reverse1(x int) int {
	var res int
	for x != 0 {
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		temp := x % 10
		x /= 10
		res = res*10 + temp
	}
	return res
}
