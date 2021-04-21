package is_unique_lcci

// 前后字符下标所指向的值对比
func isUnique(astr string) bool {
	res := true
	var i int
	for i = 0; i < len(astr); i++ {
		for j := len(astr) - 1; j > i; j-- {
			if astr[i] == astr[j] {
				res = false
				break
			}
		}
	}
	return res
}

// 利用比特位去实现
func isUnique1(astr string) bool {
	res := true
	base := []rune("a")
	astrR := []rune(astr)
	mark := 0
	for _, v := range astrR {
		if (mark & (1 << (v - base[0]))) != 0 {
			res = true
		}
		mark = mark | (1 << (v - base[0]))
	}

	return res
}
