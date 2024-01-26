package max_sliding_windows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow1(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expect := []int{3, 3, 5, 5, 6, 7}

	actual := maxSlidingWindow1(nums, k)
	assert.Equal(t, expect, actual, "maxSlidingWindow1 execute error")
}

func TestMaxSlidingWindow2(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expect := []int{3, 3, 5, 5, 6, 7}
	actual := maxSlidingWindow2(nums, k)
	assert.Equal(t, expect, actual, "maxSlidingWindowW execute error")
}
