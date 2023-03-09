package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	s := "aaa"
	q := "a*"
	expect := true
	actual := isMatch(s, q)
	assert.Equal(t, expect, actual, "isMatch execute failed")
}
