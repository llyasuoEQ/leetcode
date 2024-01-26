package primutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected [][]int
	}{
		{
			Nums: []int{1, 2, 3},
			Expected: [][]int{
				{
					1, 2, 3,
				},
				{
					1, 3, 2,
				},
				{
					2, 1, 3,
				},
				{
					2, 3, 1,
				},
				{
					3, 1, 2,
				},
				{
					3, 2, 1,
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual := permute(testCase.Nums)
		assert.Equal(t, testCase.Expected, actual,
			"permute execute error: Nums[%v], actual[%d], expected[%d]", testCase.Nums, actual, testCase.Expected)
	}
}
