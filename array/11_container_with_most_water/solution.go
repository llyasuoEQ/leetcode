package container_with_most_water

// 方法一:
// 暴力求解，时间复杂度大，当height达到一定的数量级的时候就会超出时间的限制
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func maxArea1(height []int) int {
	i, j := 0, 0
	var res int
	for i = 0; i < len(height)-1; i++ {
		for j = i + 1; j < len(height); j++ {
			volume := (j - i) * min(height[i], height[j])
			res = max(res, volume)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二：
// 前后指针法
// 时间复杂度：O(N)
// 空间复杂度：0(1)
func maxArea2(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for j > i {
		temp := min(height[i], height[j]) * (j - i)
		res = max(res, temp)
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
