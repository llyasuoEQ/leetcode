package permutation_sequence

import "strconv"

// getPermutation 排列序列
// 解题思路：利用数学方法依次确认第k个排列的每一位是什么，然后输出，核心思路如下
// 例如：1,2,⋯,n，我们首先确定排列中的首个元素a1
// 以1为a1的排列有(n-1)！
// 以2为a1的排列有(n-1)！
// ....
// 以n为a1的排列有(n-1)！
// 那么就得处如下结论
// 如果：k < (n-1)!，那么首位为1
// 如果 (n−1)! < k ≤ 2⋅(n−1)!，那么首位为2
// ....
// 如果(n−1)⋅(n−1)! < k ≤ n⋅(n−1)!，那么首位就为n
// 当a1确定好后，如何确定a2呢
// 同样的思路，那么除去a1的值，那么就有(n-2)!个排列
func getPermutation(n int, k int) string {
	// 存储需要排列的数字
	nums := make([]int, n)
	// 存放阶乘结果，为什么是n+1，下标几表示几的阶乘
	factorial := make([]int, n+1)
	factorial[0] = 1 // 初始化，方便后面计算

	// 初始化nums数组和factorial数组
	for i := 1; i <= n; i++ {
		nums[i-1] = i
		factorial[i] = factorial[i-1] * i
	}

	k--
	result := ""
	// 依次计算每一位值
	for i := n; i >= 1; i-- {
		index := k / factorial[i-1]
		k %= factorial[i-1]
		result += strconv.Itoa(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
	}

	return result
}
