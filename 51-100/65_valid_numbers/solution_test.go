package valid_numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	testCases := []struct {
		S        string
		Expected bool
	}{
		{
			S:        "0",
			Expected: true,
		},
		{
			S:        "-1.4e12",
			Expected: true,
		},
	}
	for _, testCase := range testCases {
		actual := isNumber(testCase.S)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("isNumber execute error: s[%v], expected[%v], actual[%v]",
			testCase.S, testCase.Expected, actual))
	}
}
