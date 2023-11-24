package trapping_rain_water

// trap 接雨水
// 解法一：暴力法，时间复杂度O(n^2)
// 依次遍历每一个柱子，每个柱子的接水取决于左边最高的柱子和右边最高的柱子的最小值
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var result int
	// 遍历第一遍，计算每一列能接多少雨水
	for i, col1 := range height {
		var (
			leftMax  = 0
			rightMax = 0
		)
		// 遍历第二遍，查找当前列左边最大的柱子和右边最大柱子
		for j, col2 := range height {
			if i == j {
				continue
			}
			if j < i { // 左边最大
				leftMax = max(leftMax, col2)
			} else {
				rightMax = max(rightMax, col2)
			}
		}
		// 计算当前列能接多少雨水
		if leftMax <= col1 || rightMax <= col1 { // 表示左右边最大列都小于当前列，这样就会漏水，无法接雨水
			continue
		}
		result += (min(leftMax, rightMax) - col1) * 1 // 乘以1表示宽度为1，面积：长 * 宽
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// trap2 接雨水
// 解法二：暴力解决方法优化，左右指针法，时间复杂度：O(N)
// 解法一每次遍历柱子的时候都需要再遍历一遍柱子来计算当前柱子左右的最大的柱子，
// 相当于重复计算了，我们使用空间换时间，用maxLeft和maxRight俩个数组来记录
// 左右最大柱子
func trap2(height []int) int {
	hLen := len(height)
	if len(height) < 3 {
		return 0
	}
	var (
		result   int
		maxLeft  = make([]int, hLen)
		maxRight = make([]int, hLen)
	)
	// 计算并记录每个柱子左边柱子最大高度
	maxLeft[0] = height[0]
	for i := 1; i < hLen; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}

	// 计算并记录每个柱子右边最大的高度
	maxRight[hLen-1] = height[hLen-1]
	for j := hLen - 2; j >= 0; j-- {
		maxRight[j] = max(height[j], maxRight[j+1])
	}

	// 计算每一列能接的雨水量
	for i, col := range height {
		if col < maxLeft[i] && col < maxRight[i] {
			result += (min(maxLeft[i], maxRight[i]) - col) * 1
		}
	}
	return result
}

// trap3 接雨水
// 解法三：利用单调栈的思路解决
// 核心思路：创建单调递减栈，当遍历的当前元素大于栈顶时就形成凹槽可以接水
func trap3(height []int) int {
	hLen := len(height)
	if hLen < 1 {
		return 0
	}
	var (
		result int
		stack  []int // 创建单调递减栈，栈里存放的是元素的下标
	)
	// {4, 2, 0, 3, 2, 5}
	for i := 0; i < hLen; i++ {
		// 形成凹槽，计算凹槽能接多骚雨水
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 栈顶元素出栈

			if len(stack) == 0 { // 栈为空，无法形成凹槽
				break
			}

			// 计算能接雨水量
			width := i - stack[len(stack)-1] - 1
			h := min(height[i], height[stack[len(stack)-1]]) - height[top]
			result += width * h
		}

		// 没有形成凹槽，压栈
		stack = append(stack, i)
	}

	return result
}
