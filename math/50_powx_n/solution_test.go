package powx_n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMypow(t *testing.T) {
	testCases := []struct {
		X        float64
		N        int
		Expected float64
	}{
		{
			X:        2,
			N:        10,
			Expected: 1024,
		},
	}
	for _, testCase := range testCases {
		actual := myPow(testCase.X, testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("divide execute error: x[%f], n[%d], expected[%f], actual[%f]",
			testCase.X, testCase.N, testCase.Expected, actual))
		actual2 := myPow2(testCase.X, testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("divide execute error: x[%f], n[%d], expected[%f], actual[%f]",
			testCase.X, testCase.N, testCase.Expected, actual2))
	}
}
