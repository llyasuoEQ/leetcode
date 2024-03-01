package slimbing_stairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			N:        2,
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		actual := climbStairs(testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("climbStairs execute error: n[%v], expected[%v], actual[%v]",
			testCase.N, testCase.Expected, actual))
	}
}
