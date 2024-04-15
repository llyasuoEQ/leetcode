package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	s := "aacabdkacaa"
	expect := "aca"
	actual := longestPalindrome1(s)
	assert.Equal(t, expect, actual, "longestPalindrome1 execute error")
	actual2 := longestPalindrome2(s)
	assert.Equal(t, expect, actual2, "longestPalindrome2 execute error")
}
