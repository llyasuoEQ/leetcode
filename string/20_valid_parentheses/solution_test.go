package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()[]{}"), "isValid execute failed")
}
