package spiral_matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		Matrix   [][]int
		Expected []int
	}{
		{
			Matrix: [][]int{
				{
					1, 2, 3,
				},
				{
					4, 5, 6,
				},
				{
					7, 8, 9,
				},
			},
			Expected: []int{
				1, 2, 3, 6, 9, 8, 7, 4, 5,
			},
		},
	}
	for _, testCase := range testCases {
		actual := spiralOrder(testCase.Matrix)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("spiralOrder execute error: matrix[%v], expected[%v], actual[%v]",
			testCase.Matrix, testCase.Expected, actual))
	}
}
