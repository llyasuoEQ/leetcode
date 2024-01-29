package n_queens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	testCases := []struct {
		N        int
		Expected [][]string
	}{
		{
			N: 4,
			Expected: [][]string{
				{
					"..Q.", "Q...", "...Q", ".Q..",
				},
				{
					".Q..", "...Q", "Q...", "..Q.",
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual := solveNQueens(testCase.N)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("solveNQueens execute error: n[%d], expected[%v], actual[%v]",
			testCase.N, testCase.Expected, actual))
	}
}
