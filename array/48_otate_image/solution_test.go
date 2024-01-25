package otate_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		Matrix   [][]int
		Expected [][]int
	}{
		{
			Matrix: [][]int{
				{
					5, 1, 9, 11,
				},
				{
					2, 4, 8, 10,
				},
				{
					13, 3, 6, 7,
				},
				{
					15, 14, 12, 16,
				},
			},
			Expected: [][]int{
				{
					15, 13, 2, 5,
				},
				{
					14, 3, 4, 1,
				},
				{
					12, 6, 8, 9,
				},
				{
					16, 7, 10, 11,
				},
			},
		},
	}
	for _, testCase := range testCases {
		rotate(testCase.Matrix)
		assert.Equal(t, testCase.Expected, testCase.Matrix,
			"rotate execute error: Matrix[%v], expected[%v]", testCase.Matrix, testCase.Expected)
	}
}
