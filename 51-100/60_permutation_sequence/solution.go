package permutation_sequence

// getPermutation 排列序列
// 解题思路：利用数学方法依次确认第k个排列的每一位是什么，然后输出，核心思路如下
// 例如：1,2,⋯,n，我们首先确定排列中的首个元素a
// 以1为a的排列有(n-1)！
// 以2为a的排列有(n-1)！
// ....
// 以n为a的排列有(n-1)！
// 那么就得处如下结论
// 如果：k < (n-1)!，那么首位为1
// 如果 (n−1)!< k ≤ 2⋅(n−1)!，那么首位为2
// ....
// 如果(n−1)⋅(n−1)!<k≤n⋅(n−1)!，那么首位就为n
func getPermutation(n int, k int) string {
	// 存储需要排列的数字
	nums := make([]int, n)
	//

	return ""
}
