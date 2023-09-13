package valid_sudoku

// isValidSudoku1 直接暴力求解，一个条件一个条件检验 行、列、3X3
func isValidSudoku1(board [][]byte) bool {
	// 检测行
	for _, rows := range board {
		if !validSudokuArray(rows) {
			return false
		}
	}

	// 检测行
	for i := 0; i < 9; i++ { // 固定是9列
		var cols []byte
		for _, rows := range board {
			cols = append(cols, rows[i])
		}
		if !validSudokuArray(cols) {
			return false
		}
	}

	// 检测3x3宫格是否合规
	for row := 0; row < len(board[0])-2; row += 3 {
		for col := 0; col < len(board)-2; col += 3 {
			var cellArray []byte
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					cellArray = append(cellArray, board[row+i][col+j])
				}
			}
			if !validSudokuArray(cellArray) {
				return false
			}
		}
	}

	return true
}

func validSudokuArray(arr []byte) bool {
	valueCheckMap := make(map[byte]bool)
	for _, item := range arr {
		if item == '.' {
			continue
		}
		if valueCheckMap[item] {
			return false
		}
		valueCheckMap[item] = true
	}
	return true
}

// isValidSudoku2 先建立校验的映射，然后一次判断值是否满足要求，遍历一遍矩阵完成
func isValidSudoku2(board [][]byte) bool {
	var ( // 构建检查各行、各列、各3x3宫格的映射
		rowCheck  = make([]map[byte]bool, 9) // 行校验map
		colCheck  = make([]map[byte]bool, 9) // 列校验map
		cellCheck = make([]map[byte]bool, 9) // 校验3X3子宫格map
	)
	for row := 0; row < len(board[0]); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == '.' {
				continue
			}
			if rowCheck[col] == nil {
				rowCheck[col] = make(map[byte]bool, 9)
			}
			if colCheck[row] == nil {
				colCheck[row] = make(map[byte]bool, 9)
			}
			cellIndex := (row / 3 * 3) + (col / 3)
			if cellCheck[cellIndex] == nil {
				cellCheck[cellIndex] = make(map[byte]bool, 9)
			}
			if rowCheck[col][board[row][col]] || colCheck[row][board[row][col]] || cellCheck[cellIndex][board[row][col]] {
				return false
			}
			rowCheck[col][board[row][col]] = true
			colCheck[row][board[row][col]] = true
			cellCheck[cellIndex][board[row][col]] = true
		}
	}
	return true
}
