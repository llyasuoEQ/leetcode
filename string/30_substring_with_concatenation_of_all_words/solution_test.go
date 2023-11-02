package substring_with_concatenation_of_all_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{
		"foo",
		"bar",
	}
	expect := []int{0, 9}
	actual1 := findSubstring(s, words)
	assert.ElementsMatch(t, expect, actual1, "findSubstring execute failed")
	actual2 := findSubstring2(s, words)
	assert.ElementsMatch(t, expect, actual2, "findSubstring2 execute failed")
}
