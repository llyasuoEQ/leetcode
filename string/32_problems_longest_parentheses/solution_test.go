package problems_longest_parentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	S        string
	Expected int
}

func TestLongestValidParentheses(t *testing.T) {
	testCases := []TestCase{
		{
			S:        "(()",
			Expected: 2,
		},
	}
	for _, c := range testCases {
		actual := longestValidParentheses(c.S)
		assert.Equal(t, c.Expected, actual, fmt.Sprintf("longestValidParentheses execute error: S[%s]", c.S))
	}
}
