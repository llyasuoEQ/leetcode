package plus_one

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		Digits   []int
		Expected []int
	}{
		{
			Digits: []int{
				1, 2, 3,
			},
			Expected: []int{
				1, 2, 4,
			},
		},
	}
	for _, testCase := range testCases {
		actual := plusOne(testCase.Digits)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("plusOne execute error: digits[%v], expected[%v], actual[%v]",
			testCase.Digits, testCase.Expected, actual))
	}
}
