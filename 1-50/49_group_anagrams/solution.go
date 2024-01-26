package group_anagrams

import "sort"

// groupAnagrams 字母异位词分组
// 解题思路：使用哈希映射来将异位词分组
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// 将字符串转换为字符串数组，并进行排序
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})

		// 将排序后的字符数组转换为字符串并做为异位词标识
		sortedStr := string(strBytes)

		// 将当前字符串添加到对应的异位词分组中
		groups[sortedStr] = append(groups[sortedStr], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
