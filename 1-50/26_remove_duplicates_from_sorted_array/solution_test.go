package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	actual := removeDuplicates2(a)
	expect := 5
	assert.Equal(t, expect, actual, "removeDuplicates execute failed")
}
