package n_queens_ii

// totalNQueens N皇后II
// 解题思路：回溯算法解答，和N皇后一样的解法，只是把输出的棋盘换成输出棋盘的数量
func totalNQueens(n int) int {
	var count int
	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		rows := make([]byte, n)
		for j := 0; j < n; j++ {
			rows[j] = '.'
		}
		board[i] = rows
	}

	// 回溯计算
	backtrack(board, 0, &count)

	return count
}

// backtrack 回溯算法
func backtrack(board [][]byte, col int, count *int) {
	n := len(board)

	// 递归结束条件，所有放置完毕返回
	if col == n {
		*count++
		return
	}

	// 遍历当前行的每一列
	for row := 0; row < n; row++ {
		if isSafe(board, row, col) {
			board[row][col] = 'Q'
			backtrack(board, col+1, count)
			board[row][col] = '.'
		}
	}

}

// isSafe 是否可以放置皇后牌
// 上述回溯算法是按照每一列来依次放置的，那么只需要检查行，左上方，左下方是否有皇后牌即可
func isSafe(board [][]byte, row, col int) bool {
	n := len(board)

	// 检查行是否可以放置皇后牌
	for i := 0; i < n; i++ {
		if board[row][i] == 'Q' {
			return false
		}
	}

	// 检查左上方是否可以放置皇后牌
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检查左下方是否可以放置皇后牌
	for i, j := row+1, col-1; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
