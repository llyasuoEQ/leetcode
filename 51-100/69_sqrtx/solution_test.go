package sqrtx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		X        int
		Expected int
	}{
		{
			X:        8,
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		actual := mySqrt(testCase.X)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("mySqrt execute error: x[%v], expected[%v], actual[%v]",
			testCase.X, testCase.Expected, actual))
	}
}
