package median_of_tw_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	actual := findMedianSortedArrays2(nums1, nums2)
	expect := 2.5
	assert.Equal(t, expect, actual, "findMedianSortedArrays2 execute error")
}
