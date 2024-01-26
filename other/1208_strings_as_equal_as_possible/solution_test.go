package strings_as_equal_as_possible

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualSubstring(t *testing.T) {
	a := "abcd"
	b := "bcdf"
	maxCost := 3

	expect := 3
	actual := equalSubstring(a, b, maxCost)
	assert.Equal(t, expect, actual, "equalSubstring execute error")
}

type A struct {
	M int64
}

func (a *A) Add1() {
	a.M++
}

func (a A) Add2() {
	a.M++
}

func (a A) Println() {
	fmt.Println("a is ", a.M)
}

func (a A) AP() {
	a.M++
	a.Println()
}

func TestStruct(t *testing.T) {
	a := A{}
	a.Println()
	a.Add1()
	a.Println()
	(&a).Add1()
	a.Println()
	a.Add2()
	a.Println()
	a.AP()
}
