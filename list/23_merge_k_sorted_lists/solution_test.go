package merge_k_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	list1 := newListNode(1, newListNode(4, newListNode(5, nil)))
	list2 := newListNode(1, newListNode(3, newListNode(4, nil)))
	list3 := newListNode(2, newListNode(6, nil))
	var lists []*ListNode
	lists = append(lists, list1, list2, list3)
	expect := newListNode(1, newListNode(1, newListNode(2,
		newListNode(3, newListNode(4, newListNode(4,
			newListNode(5, newListNode(6, nil))))))))
	actual := mergeKLists(lists)
	assert.Equal(t, expect, actual, "mergeKLists execute error")
}
