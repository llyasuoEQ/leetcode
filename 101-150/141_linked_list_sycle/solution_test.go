package linked_list_sycle

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
		Expected bool
	}{
		{
			Head:     node,
			Expected: true,
		},
	}
	for _, testCase := range testCases {
		actual := hasCycle(testCase.Head)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("hasCycle execute error: node[%v], expected[%v], actual[%v]",
			testCase.Head, testCase.Expected, actual))
	}
}
