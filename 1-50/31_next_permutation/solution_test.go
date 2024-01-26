package next_permutation

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNextPermutation(t *testing.T) {
	// actual := []int{1, 1, 2}
	actual := []int{1, 3, 2}
	nextPermutation(actual)
	// expect := []int{1, 2, 1}
	expect := []int{2, 1, 3}
	assert.Equal(t, expect, actual, "nextPermutation execute error")
}
