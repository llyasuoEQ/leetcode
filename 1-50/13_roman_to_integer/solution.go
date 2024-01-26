package roman_to_integer

var romaToIntMap = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

// romanToInt 遍历求解
// 1. 枚举罗马字符和数字的映射关系
// 2. 遍历罗马字符串，依次累加遍历罗马字符对应的数字，如果当前罗马字符值大于前一位罗马字符值，减去 2 * (前置罗马字符值)
func romanToInt(s string) int {
	var (
		res  int
		prev int = 0
	)

	for i := 0; i < len(s); i++ {
		res += romaToIntMap[s[i]]
		if i > 0 && prev < romaToIntMap[s[i]] { // 此时需要考虑400,900等类似情况
			res -= prev * 2
		}
		prev = romaToIntMap[s[i]]
	}
	return res
}
