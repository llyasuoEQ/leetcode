package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	a := " "
	expect := 0
	actual := lengthOfLongestSubstring(a)
	assert.Equal(t, expect, actual, "lengthOfLongestSubstring execute error")
}
