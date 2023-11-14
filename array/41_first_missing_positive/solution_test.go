package first_missing_positive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{1, 2, 0},
			Expected: 3,
		},
		{
			Nums:     []int{3, 4, -1, 1},
			Expected: 2,
		},
		{
			Nums:     []int{7, 8, 9, 11, 12},
			Expected: 1,
		},
	}
	for _, testCase := range testCases {
		actual := firstMissingPositive(testCase.Nums)
		assert.Equal(t, testCase.Expected, actual,
			"firstMissingPositive execute error: nums[%v], acutal[%d], expected[%d]", testCase.Nums, actual, testCase.Expected)
	}
}
