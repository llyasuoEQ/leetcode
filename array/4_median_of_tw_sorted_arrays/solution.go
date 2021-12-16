package median_of_tw_sorted_arrays

// 方法一：
// 合并数组找出中位数
// 时间复杂度：O(M+N)
// 空间复杂度：O(M+N）
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var res float64
	mNums := mergeArrages(nums1, nums2)
	mLen := len(mNums)
	if mLen == 1 {
		res = float64(mNums[0])
	}
	if mLen > 1 {
		median := mLen / 2
		if mLen%2 == 0 {
			// 偶数个
			res = (float64(mNums[median]) + float64(mNums[median-1])) / 2
		} else {
			// 奇数个
			res = float64(mNums[median])
		}
	}
	return res
}

// mergeArrays...
func mergeArrages(nums1 []int, nums2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}
	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}

// 方法二：
// 二分查找法
// 时间复杂度：O(log(m+n))
// 空间复杂度：O(1)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var res float64
	totalLen := len(nums1) + len(nums2)
	k := totalLen / 2
	if totalLen%2 == 0 {
		// 偶数取俩个中位数求平均值
		res = (float64(getSortIndexKey(nums1, nums2, k)) + float64(getSortIndexKey(nums1, nums2, k+1))) / 2
	} else {
		res = float64(getSortIndexKey(nums1, nums2, k+1))
	}
	return res
}

// getSortIndexKey...
func getSortIndexKey(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		nIndex1 := min(half+index1, len(nums1)) - 1
		nIndex2 := min(half+index2, len(nums2)) - 1
		p1, p2 := nums1[nIndex1], nums2[nIndex2]
		if p1 <= p2 {
			k -= nIndex1 - index1 + 1
			index1 = nIndex1 + 1
		} else {
			k -= nIndex2 - index2 + 1
			index2 = nIndex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法三：
// 划分数组
// 时间复杂度：O(log(m+n))
// 空间复杂度：O(1)
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	return 0
}
