package jump_game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected bool
	}{
		{
			Nums: []int{
				2, 3, 1, 1, 4,
			},
			Expected: true,
		},
	}
	for _, testCase := range testCases {
		actual := canJump(testCase.Nums)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("canJump execute error: Nums[%v], expected[%v], actual[%v]",
			testCase.Nums, testCase.Expected, actual))
	}
}
