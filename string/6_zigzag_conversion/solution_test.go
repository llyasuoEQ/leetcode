package zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	s := "AB"
	numRows := 1
	actual := convert(s, numRows)
	expect := "AB"
	assert.Equal(t, expect, actual, "execute error")
}
