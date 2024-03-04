package squares_of_a_sorted_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected []int
	}{
		{
			Nums:     []int{-4, -1, 0, 3, 10},
			Expected: []int{0, 1, 9, 16, 100},
		},
	}
	for _, testCase := range testCases {
		actual := sortedSquares(testCase.Nums)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("sortedSquares execute error: nums[%v], expected[%v], actual[%v]",
			testCase.Nums, testCase.Expected, actual))
	}
}
