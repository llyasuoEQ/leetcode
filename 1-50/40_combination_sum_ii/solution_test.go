package combination_sum_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	actual := combinationSum2(candidates, target)
	expect := [][]int{
		{1, 2, 2},
		{5},
	}
	assert.Equal(t, expect, actual, "combinationSum2 execute error")
}
