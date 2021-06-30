package four_sum

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	// a := []int{1, 0, -1, 0, -2, 2}
	// a := []int{-2, -1, -1, 1, 1, 2, 2}
	a := []int{5, 5, 3, 5, 1, -5, 1, -2}
	actual := fourSum(a, 4)
	fmt.Println(actual)
}
