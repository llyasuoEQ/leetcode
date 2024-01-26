package wildcard_matching

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		S        string
		P        string
		Expected bool
	}{
		{
			S:        "aa",
			P:        "a",
			Expected: false,
		},
		{
			S:        "ab",
			P:        "*",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "?c",
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		actual := isMatch(testCase.S, testCase.P)
		assert.Equal(t, testCase.Expected, actual,
			fmt.Sprintf("isMatch execute failed: S[%s], P[%s], expected[%v], actual[%v]",
				testCase.S, testCase.P, testCase.Expected, actual))
	}
}
