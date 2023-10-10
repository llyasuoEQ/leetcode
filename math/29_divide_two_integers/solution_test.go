package divide_two_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	var (
		dividend = 10
		divisor  = 3
	)
	expect := 3
	actual := divide(dividend, divisor)
	assert.Equal(t, expect, actual, "divide execute error")
}
