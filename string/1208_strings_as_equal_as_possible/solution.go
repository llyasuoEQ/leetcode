package strings_as_equal_as_possible

import "math"

func equalSubstring(s string, t string, maxCost int) int {
	sr := []rune(s)
	tr := []rune(t)
	var left, right, sum, res int
	for right < len(s) {
		sum += int(math.Abs(float64(sr[right] - tr[right])))
		right++
		for sum > maxCost {
			sum -= int(math.Abs(float64(sr[left] - tr[left])))
			left++
		}
		res = int(math.Max(float64(res), float64(right-left)))
	}

	return res
}
