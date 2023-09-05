package combination_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	actual := combinationSum(candidates, target)
	expect := [][]int{
		{
			2, 2, 3,
		},
		{
			7,
		},
	}
	assert.Equal(t, expect, actual, "combinationSum execute error")
}
