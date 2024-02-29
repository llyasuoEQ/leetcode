package add_binary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		A        string
		B        string
		Expected string
	}{
		{
			A:        "11",
			B:        "1",
			Expected: "100",
		},
	}
	for _, testCase := range testCases {
		actual := addBinary(testCase.A, testCase.B)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("addBinary execute error: a[%v], b[%v], expected[%v], actual[%v]",
			testCase.A, testCase.B, testCase.Expected, actual))
	}
}
