package count_good_meals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPairs(t *testing.T) {
	nums := []int{32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
	expect := 45
	actual1 := countPairs1(nums)
	actual2 := countPairs2(nums)
	assert.Equal(t, expect, actual1, "actual1 execute error")
	assert.Equal(t, expect, actual2, "actual2 execute error")
}
