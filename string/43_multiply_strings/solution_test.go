package multiply_strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	fmt.Println()
	testCases := []struct {
		Num1     string
		Num2     string
		Expected string
	}{
		{
			Num1:     "2",
			Num2:     "3",
			Expected: "6",
		},
	}
	for _, testCase := range testCases {
		actual := multiply(testCase.Num1, testCase.Num2)
		assert.Equal(t, testCase.Expected, actual,
			fmt.Sprintf("multiply execute failed: Num1[%s], Num2[%s], expected[%s], actual[%s]",
				testCase.Num1, testCase.Num2, testCase.Expected, actual))
	}
}
