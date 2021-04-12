package strings_as_equal_as_possible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualSubstring(t *testing.T) {
	a := "abcd"
	b := "bcdf"
	maxCost := 3

	expect := 3
	actual := equalSubstring(a, b, maxCost)
	assert.Equal(t, expect, actual, "equalSubstring execute error")
}
