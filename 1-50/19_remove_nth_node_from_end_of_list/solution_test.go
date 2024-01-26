package removen_th_node_from_end_of_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := NewListNode(1, NewListNode(2, NewListNode(3, nil)))
	expect := NewListNode(1, NewListNode(2, nil))
	assert.Equal(t, expect, removeNthFromEnd(head, 2), "removeNthFromEnd execute failed")
	head1 := NewListNode(1, NewListNode(2, nil))
	expect1 := NewListNode(2, nil)
	assert.Equal(t, expect1, removeNthFromEnd(head1, 1), "removeNthFromEnd2 execute failed")
}
