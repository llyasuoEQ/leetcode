package binary_search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			Nums:     []int{-1, 0, 3, 5, 9, 12},
			Target:   9,
			Expected: 4,
		},
	}
	for _, testCase := range testCases {
		actual := search(testCase.Nums, testCase.Target)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("search execute error: nums[%v], target[%v], expected[%v], actual[%v]",
			testCase.Nums, testCase.Target, testCase.Expected, actual))
	}
}
