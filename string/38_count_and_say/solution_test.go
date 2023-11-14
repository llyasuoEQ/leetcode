package count_and_say

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	testCases := []struct {
		N        int
		Expected string
	}{
		{
			N:        1,
			Expected: "1",
		},
		{
			N:        2,
			Expected: "11",
		},
		{
			N:        3,
			Expected: "21",
		},
	}
	for _, testCase := range testCases {
		actual := countAndSay(testCase.N)
		assert.Equal(t, testCase.Expected, actual,
			"countAndSay execute failed: N[%d], expected[%s], actual[%s]", testCase.Expected, actual)
		actual2 := countAndSay2(testCase.N)
		assert.Equal(t, testCase.Expected, actual2,
			"countAndSay2 execute failed: N[%d], expected[%s], actual[%s]", testCase.Expected, actual)
	}
}
