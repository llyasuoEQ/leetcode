package n_queens_ii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTtotalNQueens(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			N:        4,
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		actual := totalNQueens(testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("totalNQueens execute error: n[%d], expected[%d], actual[%v]",
			testCase.N, testCase.Expected, actual))
	}
}
