package next_permutation

/* 在了解这个算法之前，得清楚什么是字典序:字典序，就是按照字典中出现的先后顺序进行排序
多个字符的情况：对于任意两个序列 (a,b) 和 (a’,b’)，字典序定义为： (a,b) ≤ (a′,b′) 当且仅当 a < a′ 或 (a = a′ 且 b ≤ b′).
这个题也是再考字典序算法：
给定数字序列重新排列成字典序中下一个更大的排列
以数字序列 [1,2,3]为例，其排列按照字典序依次为：
[1,2,3]
[1,3,2]
[2,1,3]
[2,3,1]
[3,1,2]
[3,2,1]
这样，排列 [2,3,1]的下一个排列即为 [3,1,2]特别的，最大的排列 [3,2,1]的下一个排列为最小的排列 [1,2,3]。
*/
// 方法一：按照字典序算法直接去找到下一个排列
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			ci := i // 交换的位置和值大小
			for j := length - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					ci = j
					break
				}
			}
			nums[ci], nums[i-1] = nums[i-1], nums[ci]
			break
		}
	}
	reverse(nums[i:])
	return
}

// [3,2,1]
func reverse(nums []int) {
	l := 0
	r := len(nums) - 1
	for r > l {
		nums[r], nums[l] = nums[l], nums[r]
		r--
		l++
	}
}
