package minimum_size_subarray_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		Target   int
		Nums     []int
		Expected int
	}{
		{
			Target:   7,
			Nums:     []int{2, 3, 1, 2, 4, 3},
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		actual := minSubArrayLen(testCase.Target, testCase.Nums)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("minSubArrayLen execute error: target[%v], nums[%v], expected[%v], actual[%v]",
			testCase.Target, testCase.Nums, testCase.Expected, actual))
	}
}
