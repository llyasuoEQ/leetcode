package two_sum

// 方法一:
// 暴力求解
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	var res []int
	for i, m := range nums {
		for j, n := range nums {
			if i != j && (m+n) == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

// 方法二：
// 前后指针的方式
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func twoSum2(nums []int, target int) []int {
	var res []int
	i := 0
	for ; i < len(nums)-1; i++ {
		j := len(nums) - 1
		for j > i {
			if (nums[i] + nums[j]) == target {
				res = append(res, i, j)
				return res
			}
			j--
		}
	}
	return res
}

// 方法三：
// 利用哈希的思路去实现
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func twoSum3(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{p, i}
		}
		hashTable[v] = i
	}
	return []int{}
}
