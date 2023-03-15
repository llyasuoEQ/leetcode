package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	x := 121
	expect := true
	actual := isPalindrome(x)
	assert.Equal(t, expect, actual, "isPalindrome execute failed")
}
