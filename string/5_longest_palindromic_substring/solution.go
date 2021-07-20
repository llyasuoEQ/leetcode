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
