package maximum_subarray

// maxSubArray 最大子数和
// 解题思路：动态规划
// dp[i]: 表示以nums[i]为结尾的连续子数组的最大和，这里为什么是以nums[i]结尾，为了保证连续性，
// 同时满足比较每个以nums[i]结尾的连续子序和找出最大和
// 当i = 0时，那么dp[0] = nums[0]
// 如果nums数组全部大于0，那么dp[i] = dp[i-1] + nums[i]
// 但是dp[i-1]可能存在负数
// - 如果dp[i-1] > 0，那么可以把nums[i]直接接在dp[i-1]表示的那个数组的后面，得到更大的连续子数组；
// - 如果dp[i-1] < 0，那么nums[i]加上前面的数dp[i-1]以后不会变大，于是单独记nums[i]的值，就是dp[i]
// - 上述俩个如果可以记为状态方程：
// - 如果 dp[i-1] > 0 dp[i] = dp[i-1] + nums[i]
// - 如果 dp[i-1] < 0 dp[i] = nums[i]
func maxSubArray(nums []int) int {
	maxSum := nums[0]     // 最大连续子序和
	currentSum := nums[0] // 当前连续子序和

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 { // 表示当前子序和小于 0
			currentSum = nums[i]
		} else {
			currentSum = currentSum + nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
