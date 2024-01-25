package otate_image

// rotate 旋转图像
// 解题思路：首先进行矩阵行转列，然后反转每一行
func rotate(matrix [][]int) {
	mLen := len(matrix)
	// 先行和列反转
	for i := 0; i < mLen; i++ {
		for j := i; j < mLen; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 然后反转每一行
	for i := 0; i < mLen; i++ {
		for j := 0; j < mLen/2; j++ {
			matrix[i][j], matrix[i][mLen-j-1] = matrix[i][mLen-j-1], matrix[i][j]
		}
	}
}
