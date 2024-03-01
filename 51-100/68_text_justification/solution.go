package text_justification

import "strings"

// fullJustify 文本左右对齐
// 解题思路：模拟放置单词
// 分为三种情况讨论
// 1. 当前行是最后一行：单词左对齐，且单词之间只有一个空格，在行末填充剩余空格；
// 2. 当前行不是最后一行，且只有一个单词：该单词左对齐，在行末填充空格；
// 3. 当前行不是最后一行，且不只一个单词：设当前行单数位numsWords，空格数位numsSpaces，我们需要将空格均匀分配在单词之间，
// 则单词之间至少有 avgSpaces = numSpaces / numWords-1 个空格，对于多出来的 extraSpaces = numSpaces % (numsWords - 1)
// 个空格，应填在前extraSpaces 个单词之间。因此，前 extraSpaces 个单词之间填充 avgSpaces+1 个空格，
// 其余单词之间填充 avgSpaces\textit{avgSpaces}avgSpaces 个空格。
func fullJustify(words []string, maxWidth int) []string {
	var (
		result    []string
		currLine  []string // 当前行数组
		currWidth int      // 当前行宽度
	)

	// 遍历单词数组
	for _, word := range words {
		// 如果当前行加上新单词的宽度和空格仍然小雨等于最大宽度，这里len(currline)表示当前行已存在的单词个数，但么就需要len(currline)个空格
		if currWidth+len(word)+len(currLine) <= maxWidth {
			currLine = append(currLine, word)
			currWidth += len(word)
		} else { // 当前行已经满了，需要进行对齐操作
			result = append(result, alignLine(currLine, currWidth, maxWidth))

			// 重置当前行的单词列表，将新单词添加到其中
			currLine = []string{word}
			// 更新当前行的宽度
			currWidth = len(word)
		}
	}

	// 处理最有一行，左对齐并补齐空格
	lastLine := strings.Join(currLine, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	result = append(result, lastLine)
	return result
}

// alignLine 行对齐
func alignLine(line []string, currWidth, maxWidth int) string {
	numWords := len(line)            // 当前行的单词数量
	numSpace := maxWidth - currWidth // 需要插入的总空格数

	// 如果只有一个单词，左对齐不齐空格
	if numWords == 1 {
		return line[0] + strings.Repeat(" ", numSpace)
	}

	// 平均每个单词之间需要补齐的空格数
	numSpaceBetween := numSpace / (numWords - 1)
	// 额外需要插入的空格数
	extraSpaces := numSpace % (numWords - 1)

	result := ""
	for i := 0; i < numWords-1; i++ {
		result += line[i]                              // 添加当前单词
		result += strings.Repeat(" ", numSpaceBetween) // 添加平均每个单词之间的空格

		if i < extraSpaces { // 如果还有额外的空格，添加一个
			result += " "
		}
	}
	result += line[numWords-1] // 添加最后一个单词
	return result
}
