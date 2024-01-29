package n_queens

// solveNQueens N皇后
// 利用回溯算法解答
//  1. 初始化一个二维棋盘
//  2. 定义backtrack进行回溯求解，其中col标识当前需要放置的皇后列号
//  3. 遍历每一行，检查当前位置是否可以放置皇后。如果可以放置那么递归进入下一列的皇后放置操作，
//     反之，遍历当前列的每一行，检查当前位置是否可以放置皇后
//  4. 当递归返回时，恢复当前位置为空，继续便利下一行
func solveNQueens(n int) [][]string {
	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		rows := make([]byte, n)
		for j := 0; j < n; j++ {
			rows[j] = '.'
		}
		board[i] = rows
	}

	// 用于存储所有解的结果
	var result [][]string
	backtrack(board, 0, &result)

	return result
}

// backtrack 回溯函数
func backtrack(board [][]byte, col int, result *[][]string) {
	// 棋盘行数
	n := len(board)

	// 递归结束条件，所有列放置完毕返回
	if col == n {
		var solution []string
		for i := 0; i < n; i++ {
			solution = append(solution, string(board[i]))
		}
		*result = append(*result, solution)
		return
	}

	// 遍历当前列的每一行
	for row := 0; row < n; row++ {
		// 检查当前的位置是否可以放置皇后
		if isSafe(board, row, col) {
			// 放置皇后
			board[row][col] = 'Q'
			// 在下一列继续放置皇后
			backtrack(board, col+1, result)
			// 恢复当前位置为空
			board[row][col] = '.'
		}
	}

}

// isSafe 检查当前位置是否可以放置皇后，因为上述是按照每一列先放置好，然后再检查下一列，所以
// 只需要检查每一行是否可以放置皇后，以及左上方和左下方，相当于检查一半即可
func isSafe(board [][]byte, row, col int) bool {
	n := len(board)

	// 检查同一行是否可以放置
	for i := 0; i < n; i++ {
		if board[row][i] == 'Q' {
			return false
		}
	}

	// 检查左上方是否可以放置
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检查左下方是否可以放置
	for i, j := row+1, col-1; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
