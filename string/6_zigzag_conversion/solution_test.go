package zigzag_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	s := "AB"
	numRows := 1
	actual := convert(s, numRows)
	expect := "AB"
	assert.Equal(t, expect, actual, "execute error")
}
