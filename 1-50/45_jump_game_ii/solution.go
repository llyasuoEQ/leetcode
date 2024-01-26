package jump_game_ii

// jump 跳跃游戏二
// 动态规划实现：dp[i]表示调到下标i的位置所需要最小跳跃次数
// 对于当前位置i，从前一个位置j(0 ≤ j < i)跳跃到位置i，需要满足：
// 位置j + nums[j] ≥ i，那么对于跳跃至i的最小跳跃次数需要满足上述
// 条件跳跃至j 加 1，即dp[i] = min(dp[i], dp[j]+1)
// 初始条件：dp[0] = 0
// 时间复杂度：O(N^2)，时间复杂度较高
func jump(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen)
	// dp每个位置存放最大跳跃次数
	for i := 0; i < numsLen; i++ {
		dp[i] = i
	}

	// 寻找最小跳跃次数
	for i := 0; i < numsLen; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[numsLen-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// jump2 跳跃游戏二
// 贪心算法：如果i为下次跳跃的起点，那么i能跳到的位置就是[i+1,i+num[i]]的左闭右闭区间
// 那么我们是i作为下次起点尽可能跳的远。
func jump2(nums []int) int {
	// 维护三个变量
	var (
		end     int // 当前能到达的最大位置
		maxPos  int // 下一步所能跳跃的最大位置
		steps   int // 最少跳跃次数
		numsLen = len(nums)
	)

	for i := 0; i < numsLen-1; i++ {
		maxPos = max(maxPos, i+nums[i]) // 计算下一步能跳跃的最大位置
		if end == i {                   // 到达边界就增加一步，且更新最大跳跃位置
			end = maxPos
			steps += 1
		}
	}
	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
