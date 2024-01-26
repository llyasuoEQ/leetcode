package remove_duplicates_from_sorted_array

// [0,1,1,1,2]
// 方法一：
// 因为是有序的，直接遍历
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func removeDuplicates1(nums []int) int {
	length := len(nums)
	for i := 1; i < length; {
		if nums[i] == nums[i-1] {
			for j := i; j < length-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			length--
			continue
		}
		i++
	}
	return length
}

// 方法二：
// 前后指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length < 1 {
		return length
	}
	slow := 1
	for fast := 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
