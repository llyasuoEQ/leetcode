package unique_paths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		M        int
		N        int
		Expected int
	}{
		{
			M:        3,
			N:        7,
			Expected: 28,
		},
	}
	for _, testCase := range testCases {
		actual := uniquePaths(testCase.M, testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("uniquePaths execute error: m[%v], n[%v], expected[%v], actual[%v]",
			testCase.M, testCase.N, testCase.Expected, actual))
	}
}
