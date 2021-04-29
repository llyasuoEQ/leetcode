package in_order_traversal_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal2(t *testing.T) {
	node1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node2 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node2.Left = &node1
	node2.Right = &node3
	expect := []int{1, 2, 3}
	actual := inorderTraversal2(&node2)
	assert.Equal(t, expect, actual, "inorderTraversal2 execute error")
}
