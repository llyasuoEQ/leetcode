package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	s := "MCMXCIV"
	expect := 1994
	assert.Equal(t, expect, romanToInt(s), "romaToInt execute failed")
}
