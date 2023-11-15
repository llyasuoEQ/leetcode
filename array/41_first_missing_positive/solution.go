package first_missing_positive

// firstMissingPositive 缺失的第一个正整数
// 思路：利用桶排序思路，一个萝卜一个坑，这里的萝卜就是数组元素
// 坑就是数组的下标，将数组元素放置再对应的下标位置，最后遍历不
// 等于下标的元素就是确实的第一个正整数
func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	// 1. 将各个元素放置对应下标的位置
	for i := 0; i < numsLen; i++ {
		for nums[i] > 0 && nums[i] <= numsLen && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 2. 找出数组元素不等于下表的元素即为缺失的第一个正整数
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return numsLen + 1
}
