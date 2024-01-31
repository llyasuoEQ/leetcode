package jump_game

// canJump 跳跃游戏
// 解题思路：使用贪心算法来解答
func canJump(nums []int) bool {
	// 标记的最大跳跃值
	var maxJump int

	// 依次遍历，利用贪心算法来寻找最优解
	for i := 0; i < len(nums); i++ {
		// 当前位置都无法跳跃到时，则表示无法跳跃到最后一个下标，直接结束
		if i > maxJump {
			return false
		}

		maxJump = max(maxJump, i+nums[i])

		// 当最大跳跃值大于数组最后一个下标时，说明就能跳跃到最后一个下标
		if maxJump >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
