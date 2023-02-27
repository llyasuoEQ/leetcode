package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	x := 123
	actual := reverse(x)
	expect := 321
	assert.Equal(t, expect, actual, "reverse execute failed")
}
