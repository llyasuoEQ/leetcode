package text_justification

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	testCases := []struct {
		Words    []string
		MaxWidth int
		Expected []string
	}{
		{
			Words: []string{
				"This", "is", "an", "example", "of", "text", "justification.",
			},
			MaxWidth: 16,
			Expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
	}
	for _, testCase := range testCases {
		actual := fullJustify(testCase.Words, testCase.MaxWidth)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("fullJustify execute error: words[%v], maxWidth[%v], expected[%v], actual[%v]",
			testCase.Words, testCase.MaxWidth, testCase.Expected, actual))
	}
}
