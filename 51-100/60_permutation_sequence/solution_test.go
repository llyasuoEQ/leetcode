package permutation_sequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		N        int
		K        int
		Expected string
	}{
		{
			N:        3,
			K:        3,
			Expected: "213",
		},
	}
	for _, testCase := range testCases {
		actual := getPermutation(testCase.N, testCase.K)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("getPermutation execute error: n[%v], k[%v], expected[%v], actual[%v]",
			testCase.N, testCase.K, testCase.Expected, actual))
	}
}
