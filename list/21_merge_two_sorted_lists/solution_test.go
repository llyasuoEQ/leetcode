package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := newListNode(1, newListNode(2, newListNode(4, nil)))
	list2 := newListNode(1, newListNode(3, newListNode(4, nil)))
	expect := newListNode(1, newListNode(1, newListNode(2, newListNode(3, newListNode(4, newListNode(4, nil))))))
	actual := mergeTwoLists(list1, list2)
	assert.Equal(t, expect, actual, "mergeTwoLists execute failed")
}
