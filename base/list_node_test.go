package base

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	head := NewListNode(1, NewListNode(2, nil))
	fmt.Println(head)
}
