package longest_common_prefix

import "math"

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
