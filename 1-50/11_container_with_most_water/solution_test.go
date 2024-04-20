package container_with_most_water

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {

	testCases := []struct {
		Height   []int
		Expected int
	}{
		{
			Height: []int{
				1, 8, 6, 2, 5, 4, 8, 3, 7,
			},
			Expected: 49,
		},
	}
	for _, testCase := range testCases {
		actual1 := maxArea1(testCase.Height)
		assert.Equal(t, testCase.Expected, actual1, fmt.Sprintf("maxArea1 execute error: height[%v], expected[%v], actual[%v]",
			testCase.Height, testCase.Expected, actual1))
		actual2 := maxArea2(testCase.Height)
		assert.Equal(t, testCase.Expected, actual2, fmt.Sprintf("maxArea2 execute error: height[%v], expected[%v], actual[%v]",
			testCase.Height, testCase.Expected, actual2))
	}
}
