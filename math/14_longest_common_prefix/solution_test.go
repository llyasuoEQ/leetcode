package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expect := "fl"
	assert.Equal(t, expect, longestCommonPrefix(strs), "longestCommonPrefix execute failed")
}
