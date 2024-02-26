package unique_paths_ii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := []struct {
		ObstacleGrid [][]int
		Expected     int
	}{
		{
			ObstacleGrid: [][]int{
				{
					0, 0, 0,
				},
				{
					0, 1, 0,
				},
				{
					0, 0, 0,
				},
			},
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		actual := uniquePathsWithObstacles(testCase.ObstacleGrid)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("uniquePathsWithObstacles execute error: obstacleGrid[%v], expected[%v], actual[%v]",
			testCase.ObstacleGrid, testCase.Expected, actual))
	}
}
