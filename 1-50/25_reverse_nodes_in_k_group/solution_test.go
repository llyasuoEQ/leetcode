package reverse_nodes_in_k_group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	head := newListNode(1, newListNode(2, newListNode(3, newListNode(4, newListNode(5, nil)))))
	k := 2
	expect := newListNode(2, newListNode(1, newListNode(4, newListNode(3, newListNode(5, nil)))))
	actual := reverseKGroup(head, k)
	assert.Equal(t, expect, actual, "reverseKGroup execute failed")
}
