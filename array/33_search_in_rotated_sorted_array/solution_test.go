package search_in_rotated_sorted_array

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSearch(t *testing.T) {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 0
	// expect := 4
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 3
	// expect := -1
	nums := []int{5, 1, 3}
	target := 5
	expect := 0
	actual := search(nums, target)
	assert.Equal(t, expect, actual, "search execute error!")
}

func BenchmarkSearch(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i := 0; i < b.N; i++ {
		search(nums, 5)
	}
}
