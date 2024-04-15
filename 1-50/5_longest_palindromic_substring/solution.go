package longest_palindromic_substring

// s = "babad"
// 方法一：左右指针法，相当于暴力求解
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func longestPalindrome1(s string) string {
	if s == "" {
		return ""
	}
	sr := []rune(s)
	rr := sr[0:1]
	length := len(sr)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if sr[j] == sr[i] {
				l, r := i, j
				for r > l {
					if sr[l] == sr[r] {
						r--
						l++
						continue
					}
					break
				}
				if r <= l && (j-i+1) > len(rr) {
					rr = sr[i : j+1]
				}
			}
		}
	}
	return string(rr)
}

// 方法二：中心扩展法
// 奇数和偶数长度俩中求职
func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 奇数长度
		len1 := expandAroudCenter(s, i, i)
		// 偶数长度
		len2 := expandAroudCenter(s, i, i+1)
		maxLen := len1
		if len1 < len2 {
			maxLen = len2
		}

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroudCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
