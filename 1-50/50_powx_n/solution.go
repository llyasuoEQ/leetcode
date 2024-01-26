package powx_n

// myPow 计算x的n次方，相当于暴力求解 O(N)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := x
	for n > 1 {
		result *= x
		n--
	}
	return result
}

// myPow2 优化上面算法，上述算法是时效太低，O(logN)
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		// 节省计算次数，相当于把result要乘的积算好
		x *= x
		n /= 2
	}
	return result
}
