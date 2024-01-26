package max_fixed_length_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualSubstring(t *testing.T) {
	a := "iib"
	k := 2

	expect := 2
	actual := maxVowels(a, k)
	assert.Equal(t, expect, actual, "equalSubstring execute error")
}
