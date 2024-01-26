package count_good_meals

import (
	"math"
)

// 方法一：暴力求解-数据源多的情况下，会超出时间限制
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func countPairs1(deliciousness []int) int {
	var res int
	length := len(deliciousness)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			sum := deliciousness[i] + deliciousness[j]
			if is2Power(sum) {
				res++
			}
		}
	}
	return res
}

func is2Power(num int) bool {
	if num == 1 {
		return true
	}
	power := 1
	for i := 1; i < math.MaxInt32; i++ {
		power = power * 2
		if power == num {
			return true
		}
		if power > num {
			return false
		}
	}
	return false
}

// 方法二：哈希表去实现
// 时间复杂度：O(NlogC)，C为最大值
// 空间复杂度：O(N)
func countPairs2(deliciousness []int) int {
	var res int
	var max int
	// 找到最大值
	for _, v := range deliciousness {
		if v > max {
			max = v
		}
	}
	maxSum := max << 1
	hashTable := map[int]int{}
	for _, v := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			res += hashTable[sum-v]
		}
		hashTable[v]++
	}
	return res % (1e9 + 7)
}
