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
	}
	for _, testCase := range testCases {
		actual := trap(testCase.Height)
		assert.Equal(t, testCase.Expected, actual,
			"trap execute error: Height[%v], acutal[%d], expected[%d]", testCase.Height, actual, testCase.Expected)
	}
}
