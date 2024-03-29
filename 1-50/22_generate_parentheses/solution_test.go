package generate_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	expect := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	acutal := generateParenthesis(n)
	assert.ElementsMatch(t, expect, acutal, "generateParenthesis execute failed")
}
