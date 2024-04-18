package post_order_traversal_of_binary_tree

import (
	"fmt"
	"myProject/leetcode/base"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := []struct {
		Root     *base.TreeNode
		Expected []int
	}{
		{
			Root: base.NewTreeNode(
				1, base.NewTreeNode(2, nil, nil), base.NewTreeNode(3, nil, nil),
			),
			Expected: []int{2, 3, 1},
		},
	}
	for _, testCase := range testCases {
		actual := postorderTraversal(testCase.Root)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("postorderTraversal execute error: root[%v], expected[%v], actual[%v]",
			testCase.Expected, testCase.Expected, actual))
		actual1 := postorderTraversal1(testCase.Root)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("postorderTraversal1 execute error: root[%v], expected[%v], actual[%v]",
			testCase.Expected, testCase.Expected, actual1))
	}

}
