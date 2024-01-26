package letter_combinations_of_a_phone_number

var digitLetter = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// letterCombinations 穷举法
func letterCombinations(digits string) []string {
	var res []string
	for i := 0; i < len(digits); i++ {
		var (
			next []string
			temp = res
		)
		if letter, ok := digitLetter[digits[i]]; ok {
			for j := 0; j < len(letter); j++ {
				next = append(next, string(letter[j]))
			}
		}
		var newTemp []string
		if len(temp) == 0 {
			newTemp = append(newTemp, next...)
		} else {
			for m := 0; m < len(temp); m++ {
				for n := 0; n < len(next); n++ {
					newTemp = append(newTemp, temp[m]+next[n])
				}
			}
		}
		res = newTemp
	}
	return res
}
