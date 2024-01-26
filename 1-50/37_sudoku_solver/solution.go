package sudoku_solver

// solveSudoku 利用回溯方法来解决
func solveSudoku(board [][]byte) {
	solve(board)
}

// solve 解独数回溯函数
func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ { // 依次遍历宫格
			if board[i][j] == '.' {
				for num := '1'; num <= '9'; num++ {
					if isValid(board, i, j, byte(num)) { // 满足条件，继续往下走
						board[i][j] = byte(num)
						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num || board[(row/3*3)+(i/3)][(col/3*3)+(i%3)] == num {
			return false
		}
	}
	return true
}
