package integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	num := 3
	expect := "III"
	actual := intToRoman(num)
	assert.Equal(t, expect, actual, "intToToman execute failed")
}
