package zigzag_conversion

// convert 利用二维矩阵填充的方法
func convert(s string, numRows int) string {
	sLen := len(s)
	if numRows < 2 || sLen <= numRows {
		return s
	}
	// 计算周期间隔
	interval := numRows*2 - 2
	// 计算二维矩阵需要多长
	mLen := (sLen + interval - 1) / interval * (numRows - 1)
	matrix := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, mLen)
	}

	x, y := 0, 0
	for i, v := range s {
		matrix[x][y] = byte(v)
		if i%interval < (numRows - 1) {
			x++
		} else {
			x--
			y++
		}
	}
	var res []byte
	for _, m := range matrix {
		for _, n := range m {
			if n > 0 {
				res = append(res, n)
			}
		}
	}
	return string(res)
}
