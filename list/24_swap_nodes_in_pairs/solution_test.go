package swap_nodes_in_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	head := newListNode(1, newListNode(2, newListNode(3, newListNode(4, nil))))
	expect := newListNode(2, newListNode(1, newListNode(4, newListNode(3, nil))))
	actual := swapPairs(head)
	assert.Equal(t, expect, actual, "swapPairs execute failed")

	head1 := newListNode(1, newListNode(2, newListNode(3, newListNode(4, nil))))
	actual1 := swapPairs1(head1)
	assert.Equal(t, expect, actual1, "swapPairs1 execute failed")
}
