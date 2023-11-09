package divide_two_integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Dividend int
	Divisor  int
	Expected int
}

func TestDivide(t *testing.T) {
	testCases := []TestCase{
		{
			Dividend: 10,
			Divisor:  3,
			Expected: 3,
		},
		{
			Dividend: 7,
			Divisor:  -3,
			Expected: -2,
		},
		{
			Dividend: 2147483647,
			Divisor:  3,
			Expected: 715827882,
		},
		{
			Dividend: -2147483648,
			Divisor:  -1,
			Expected: 2147483647,
		},
	}
	for _, c := range testCases {
		actual := divide(c.Dividend, c.Divisor)
		assert.Equal(t, c.Expected, actual, fmt.Sprintf("divide execute error: dividend[%d], divisor[%d]", c.Dividend, c.Divisor))
		actual = divide2(c.Dividend, c.Divisor)
		assert.Equal(t, c.Expected, actual, fmt.Sprintf("divide2 execute error: dividend[%d], divisor[%d]", c.Dividend, c.Divisor))
	}
}
