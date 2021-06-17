package three_sum

import (
	"fmt"
	"testing"
)

func editMap(mm map[string]string) {
	mm["aa"] = "aaaa"
}

func TestMap(t *testing.T) {
	mm := make(map[string]string)
	mm["aa"] = "bbbbb"
	fmt.Println(mm)
	editMap(mm)
	fmt.Printf("%s", mm["aa"])
	fmt.Println(mm)
}
