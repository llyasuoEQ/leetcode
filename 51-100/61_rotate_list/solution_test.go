package rotate_list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewListNode(val int, node *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: node,
	}
}

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		Head     *ListNode
		K        int
		Expected *ListNode
	}{
		{
			Head:     NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil))))),
			K:        2,
			Expected: NewListNode(4, NewListNode(5, NewListNode(1, NewListNode(2, NewListNode(3, nil))))),
		},
	}
	for _, testCase := range testCases {
		actual := rotateRight(testCase.Head, testCase.K)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("rotateRight execute error: head[%v], k[%v], expected[%v], actual[%v]",
			testCase.Head, testCase.K, testCase.Expected, actual))
	}
}
