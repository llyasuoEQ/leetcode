package find_the_index_of_the_first_occurrence_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	haystack := "mississippi"
	needle := "issip"
	expect := 4
	actual := strStr(haystack, needle)
	assert.Equal(t, expect, actual, "strStr execute error")
	actual = strStr2(haystack, needle)
	assert.Equal(t, expect, actual, "strStr2 execute error")
}
