package merge_intervals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {

	testCases := []struct {
		Intervals [][]int
		Expected  [][]int
	}{
		{
			Intervals: [][]int{
				{
					1, 3,
				},
				{
					2, 6,
				},
				{
					8, 10,
				},
				{
					15, 18,
				},
			},
			Expected: [][]int{
				{
					1, 6,
				},
				{
					8, 10,
				},
				{
					15, 18,
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual := merge(testCase.Intervals)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("merge execute error: intervals[%v], expected[%v], actual[%v]",
			testCase.Intervals, testCase.Expected, actual))
	}
}
