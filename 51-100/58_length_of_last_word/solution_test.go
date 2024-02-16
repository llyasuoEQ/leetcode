package length_of_last_word

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		S        string
		Expected int
	}{
		{
			S:        "Hello World",
			Expected: 5,
		},
	}
	for _, testCase := range testCases {
		actual := lengthOfLastWord(testCase.S)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("lengthOfLastWord execute error: s[%v], expected[%v], actual[%v]",
			testCase.S, testCase.Expected, actual))
	}
}
