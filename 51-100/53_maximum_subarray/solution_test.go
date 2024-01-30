package maximum_subarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums: []int{
				-2, 1, -3, 4, -1, 2, 1, -5, 4,
			},
			Expected: 6,
		},
	}
	for _, testCase := range testCases {
		actual := maxSubArray(testCase.Nums)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("maxSubArray execute error: nums[%v], expected[%d], actual[%v]",
			testCase.Nums, testCase.Expected, actual))
	}
}
