package jump_game_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{2, 3, 1, 1, 4},
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		actual := jump(testCase.Nums)
		assert.Equal(t, testCase.Expected, actual,
			"trap2 execute error: Nums[%v], actual[%d], expected[%d]", testCase.Nums, actual, testCase.Expected)
	}
}
