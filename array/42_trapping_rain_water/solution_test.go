package trapping_rain_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	testCases := []struct {
		Height   []int
		Expected int
	}{
		{
			Height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			Expected: 6,
		},
		{
			Height:   []int{4, 2, 0, 3, 2, 5},
			Expected: 9,
		},
	}

	for _, testCase := range testCases {
		actual := trap(testCase.Height)
		assert.Equal(t, testCase.Expected, actual,
			"trap execute error: Height[%v], acutal[%d], expected[%d]", testCase.Height, actual, testCase.Expected)
		actual2 := trap2(testCase.Height)
		assert.Equal(t, testCase.Expected, actual2,
			"trap2 execute error: Height[%v], actual2[%d], expected[%d]", testCase.Height, actual2, testCase.Expected)
		actual3 := trap3(testCase.Height)
		assert.Equal(t, testCase.Expected, actual3,
			"trap2 execute error: Height[%v], actual3[%d], expected[%d]", testCase.Height, actual3, testCase.Expected)
	}
}
