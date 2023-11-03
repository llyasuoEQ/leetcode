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

// findSubstring2
func findSubstring2(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}
	return nil
}

// func findSubstring2(s string, words []string) []int {
// 	if len(words) == 0 || len(words[0]) == 0 {
// 		return []int{}
// 	}

// 	wordLen := len(words[0])
// 	totalLen := len(words) * wordLen
// 	wordCount := make(map[string]int)

// 	for _, word := range words {
// 		wordCount[word]++
// 	}

// 	result := []int{}

// 	for i := 0; i < wordLen; i++ {
// 		left, right := i, i
// 		currentCount := make(map[string]int)

// 		for right+wordLen <= len(s) {
// 			word := s[right : right+wordLen]
// 			right += wordLen

// 			if wordCount[word] == 0 {
// 				// Reset the window and counters
// 				currentCount = make(map[string]int)
// 				left = right
// 			} else {
// 				currentCount[word]++
// 				for currentCount[word] > wordCount[word] {
// 					removedWord := s[left : left+wordLen]
// 					currentCount[removedWord]--
// 					left += wordLen
// 				}

// 				if right-left == totalLen {
// 					result = append(result, left)
// 				}
// 			}
// 		}
// 	}

// 	return result
// }
