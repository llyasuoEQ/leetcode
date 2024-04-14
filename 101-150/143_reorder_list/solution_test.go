package reorder_list

import (
	"fmt"
	"myProject/leetcode/base"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	testCases := []struct {
		Head     *base.ListNode
		Expected *base.ListNode
	}{
		{
			Head: base.NewListNode(1,
				base.NewListNode(2,
					base.NewListNode(3,
						base.NewListNode(4, nil)))),
			Expected: base.NewListNode(1,
				base.NewListNode(4,
					base.NewListNode(2,
						base.NewListNode(3, nil)))),
		},
	}
	for _, testCase := range testCases {
		// reorderList(testCase.Head)
		reordrList2(testCase.Head)
		t.Log(testCase.Head)
		assert.Equal(t, testCase.Expected, testCase.Head, fmt.Sprintf("reorderList execute error: head[%v], expected[%v]",
			testCase.Head, testCase.Expected))
	}

}
