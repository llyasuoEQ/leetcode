package find_the_index_of_the_first_occurrence_in_a_string

// strStr 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// strStr2 找出字符串中第一个匹配项的下标
func strStr2(haystack string, needle string) int {
	haystackLength := len(haystack)
	needleLength := len(needle)
	for i := 0; i < haystackLength-needleLength+1; i++ {
		var j int
		for ; j < needleLength; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == needleLength {
			return i
		}
	}

	return -1
}
