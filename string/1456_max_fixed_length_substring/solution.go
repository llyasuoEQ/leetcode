package max_fixed_length_substring

import "math"

func isVowel(c string) int {
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
		return 1
	}
	return 0
}

func maxVowels(s string, k int) int {
	sr := []rune(s)
	sLen := len(sr)
	right := 0
	sum := 0 // 记录当前窗口的值
	max := 0 // 记录窗口的最大值
	for right < sLen {
		sum += isVowel(string(sr[right]))
		right++
		if right >= k {
			max = int(math.Max(float64(max), float64(sum)))
			sum -= isVowel(string(sr[right-k]))
		}
	}
	return max
}
