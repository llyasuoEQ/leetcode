package longest_common_prefix

import (
	"math"
)

// longestCommonPrefix ...
func longestCommonPrefix(strs []string) string {
	var res []byte
	// 寻找最短的字符串长度
	var minLen = math.MaxInt32
	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	for i := 0; i < minLen; i++ {
		temp := strs[0]
		for _, v := range strs {
			if temp[i] != v[i] {
				return string(res)
			}
			temp = v
		}
		res = append(res, temp[i])
	}

	return string(res)
}

// longestCommonPrefix2 利用二分法查找法求解
func longestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	// 寻找最小的字符串
	minLen := len(strs[0])
	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	// 判断是否为子串
	isPrefix := func(length int) bool {
		temp, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != temp {
				return false
			}
		}
		return true
	}

	left, right := 0, minLen
	for left < right {
		mid := (left + right + 1) / 2
		if isPrefix(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return strs[0][:left]
}
