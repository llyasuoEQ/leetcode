package insert_interval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		Intervals   [][]int
		NewInterval []int
		Expected    [][]int
	}{
		{
			Intervals: [][]int{
				{
					1, 3,
				},
				{
					6, 9,
				},
			},
			NewInterval: []int{
				2, 5,
			},
			Expected: [][]int{
				{
					1, 5,
				},
				{
					6, 9,
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual := insert(testCase.Intervals, testCase.NewInterval)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("insert execute error: intervals[%v], newInterval[%v], expected[%v], actual[%v]",
			testCase.Intervals, testCase.NewInterval, testCase.Expected, actual))
	}
}
