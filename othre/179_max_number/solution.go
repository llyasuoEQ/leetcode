package max_number

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNumber1(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := nums[i]
		y := nums[j]
		d := 10
		// 计算x的位数
		xDight := 10
		xTemp := x / xDight
		for xTemp > 0 {
			xDight = xDight * d
			xTemp = x / xDight
		}

		// 计算y的位数
		yDight := 10
		yTemp := y / yDight
		for yTemp > 0 {
			yDight = yDight * d
			yTemp = y / yDight
		}
		return (x*yDight + y) < (y*xDight + x)
	})
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return "0"
	}
	var resBytes bytes.Buffer
	for i := len(nums) - 1; i >= 0; i-- {
		resBytes.WriteString(strconv.Itoa(nums[i]))
	}

	return resBytes.String()
}

// 第二种实现和第一种思路是一样的
// 从俩种细节方面的实现还是差距挺大的，仅供阅读
func largestNumber2(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for x >= sx {
			sx *= 10
		}
		for y >= sy {
			sy *= 10
		}
		return x*sy+y > y*sx+x
	})
	if nums[0] == 0 {
		return "0"
	}
	var resBytes bytes.Buffer
	for i := 0; i < len(nums); i++ {
		resBytes.WriteString(strconv.Itoa(nums[i]))
	}

	return resBytes.String()
}
