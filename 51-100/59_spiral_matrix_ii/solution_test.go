package spiral_matrix_ii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		N        int
		Expected [][]int
		int
	}{
		{
			N: 3,
			Expected: [][]int{
				{
					1, 2, 3,
				},
				{
					8, 9, 4,
				},
				{
					7, 6, 5,
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual := generateMatrix(testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("generateMatrix execute error: n[%v], expected[%v], actual[%v]",
			testCase.N, testCase.Expected, actual))
	}
}
