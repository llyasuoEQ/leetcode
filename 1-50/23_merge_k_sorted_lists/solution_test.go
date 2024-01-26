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
	actual1 := mergeKLists(lists)

	// 测试第二种方法
	list1 = newListNode(1, newListNode(4, newListNode(5, nil)))
	list2 = newListNode(1, newListNode(3, newListNode(4, nil)))
	list3 = newListNode(2, newListNode(6, nil))
	var list4 *ListNode
	var lists2 []*ListNode
	lists2 = append(lists2, list1, list2, list3, list4)
	actual2 := mergeKLists2(lists2)
	assert.Equal(t, expect, actual1, "mergeKLists execute error")
	assert.Equal(t, expect, actual2, "mergeKLists2 execute error")
}
