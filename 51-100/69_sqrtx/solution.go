package sqrtx

// mySqrt x的平方根
// 解题思路：使用二分查找的思想来逼近平方根
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		product := mid * mid
		if product == x {
			return mid
		} else if product < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 这里为什么返回right，因为要把小数点都抹去，所以返回right
	return right
}
