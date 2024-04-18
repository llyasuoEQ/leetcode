package preorder_traversal_binary_tree

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
			Expected: []int{1, 2, 3},
		},
	}
	for _, testCase := range testCases {
		actual := preorderTraversal(testCase.Root)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("preorderTraversal execute error: root[%v], expected[%v], actual[%v]",
			testCase.Expected, testCase.Expected, actual))
		actual1 := preorderTraversal2(testCase.Root)
		assert.Equal(t, testCase.Expected, actual, fmt.Sprintf("preorderTraversal1 execute error: root[%v], expected[%v], actual[%v]",
			testCase.Expected, testCase.Expected, actual1))
	}

}
