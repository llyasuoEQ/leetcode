package longest_substring_without_repeating_characters

// 判断重复字符
func lengthOfLongestSubstring(s string) int {
	sR := []rune(s)
	sMap := make(map[rune]int)
	slen := len(s)
	r, ans := 0, 0
	for i := 0; i < slen; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(sMap, sR[i-1])
		}
		for r < slen && sMap[sR[r]] == 0 {
			sMap[sR[r]]++
			r++
		}
		ans = max(ans, r-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 暴力求解 - 时间复杂度大
// 利用滑动窗口 从大到小，窗口中出现不重复子串时就结束
func lengthOfLongestSubstring1(s string) int {
	var res int
	if len(s) == 0 {
		return res
	}
	length := len(s)
	for i := length; i > 0; i-- {
		left, right := 0, i-1
		for right < length {
			if !helper(s[left : right+1]) {
				res = i
				return res
			}
			left++
			right++
		}
	}
	return res
}

// 判断是否有重复字符
func helper(s string) bool {
	res := false
	var i int
	for i = 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				res = true
				break
			}
		}
	}
	return res
}
