package max_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	expect := "9534330"
	actual1 := largestNumber1(nums)
	assert.Equal(t, expect, actual1, "largestNumber1 execute error")

	actual2 := largestNumber2(nums)
	assert.Equal(t, expect, actual2, "largestNumber2 execute error")
}
