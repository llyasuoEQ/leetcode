package minimum_paths_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	testCases := []struct {
		Grid     [][]int
		Expected int
	}{
		{
			Grid: [][]int{
				{
					1, 3, 1,
				},
				{
					1, 5, 1,
				},
				{
					4, 2, 1,
				},
			},
			Expected: 7,
		},
	}
	for _, testCase := range testCases {
		actual := minPathSum(testCase.Grid)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("minPathSum execute error: grid[%v], expected[%v], actual[%v]",
			testCase.Grid, testCase.Expected, actual))
	}
}
