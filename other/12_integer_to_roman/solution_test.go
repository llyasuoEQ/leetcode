package integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	num := 3
	expect := "III"
	actual1 := intToRoman(num)
	actual2 := intToRoman2(num)
	assert.Equal(t, expect, actual1, "intToToman execute failed")
	assert.Equal(t, expect, actual2, "intToToman2 execute failed")
}
