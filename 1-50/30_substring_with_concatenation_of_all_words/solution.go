package substring_with_concatenation_of_all_words

// findSubstring 实现思路如下
// 1. 计算出words中所有组合子串
// 2. 依次遍历子串，看是是否满足是子串，满足则收集该子串位于s的下标然后反回
// 这种方法时间复杂度较大，可以理解为 O(n*n)，时间复杂度较大，不能跑全case，有的会超出时间限制
func findSubstring(s string, words []string) []int {
	var result []int
	if len(words) < 1 {
		return result
	}
	substringLength := len(words) * len(words[0])
	sLength := len(s)
	if sLength < substringLength {
		return result
	}
	// 1. 计算出words中所有组合的子串
	substrings := combineStrings(words)
	// 2.  依次遍历子串，看是是否满足是子串，满足则收集该子串位于s的下标然后反回
	for _, substring := range substrings {
		substringLength := len(substring)
		for i := 0; i < (sLength - substringLength + 1); i++ {
			if s[i:i+substringLength] == substring {
				result = append(result, i)
			}
		}
	}
	return result
}

// combineStrings 同时保证组合唯一性
func combineStrings(words []string) []string {
	var (
		result []string
		used   = make(map[string]bool)
	)

	combine("", words, &result, used)
	return result
}

func combine(prefix string, words []string, result *[]string, used map[string]bool) {
	if len(words) == 0 {
		if !used[prefix] {
			*result = append(*result, prefix)
		}
		used[prefix] = true
		return
	}
	for i, word := range words {
		remainingWords := append([]string{}, words[:i]...)
		remainingWords = append(remainingWords, words[i+1:]...)
		combine(prefix+word, remainingWords, result, used)
	}
}

// findSubstring2 滑动窗口解法
func findSubstring2(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}
	wordLen := len(words[0])
	totalLen := len(words) * wordLen
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	var result []int
	// 滑动窗口的大小为单词长度，窗口每次滑动wordLen长度，所以遍历wordLen长度次
	for i := 0; i < wordLen; i++ {
		left, right := i, i
		currentCount := make(map[string]int)

		for right+wordLen < len(s) {
			word := s[right : right+wordLen]
			right = right + wordLen

			if wordCount[word] == 0 { // 不在wordCount中，不满足子串规则
				left = right
				currentCount = make(map[string]int)
			} else {
				currentCount[word]++
				// 去除重复，比如words=[af,ef], s = afafef，显然afaf不满足要去除第一个af
				for currentCount[word] > wordCount[word] {
					removeWord := s[left : left+wordLen]
					currentCount[removeWord]--
					left = left + wordLen
				}

				if (right - left) == totalLen {
					result = append(result, left)
				}
			}
		}
	}

	return result
}
