package string_to_integer_atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	a := "-91283472332"
	expect := -2147483648
	assert.Equal(t, expect, myAtoi(a), "myAtoi execute failed")
}
