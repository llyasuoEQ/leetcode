package linked_list_cycle_ii

import (
	"fmt"
	"myProject/leetcode/base"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	node := base.NewListNode(1, base.NewListNode(2, base.NewListNode(3, nil)))
	node.Next.Next.Next = node
	testCases := []struct {
		Head     *base.ListNode
		Expected *base.ListNode
	}{
		{
			Head:     node,
			Expected: node,
		},
	}
	for _, testCase := range testCases {
		actual := detectCycle(testCase.Head)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("detectCycle execute error: node[%v], expected[%v], actual[%v]",
			testCase.Head, testCase.Expected, actual))
	}
}
