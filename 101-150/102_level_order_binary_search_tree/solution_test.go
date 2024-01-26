package level_order_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
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
	node1.Left = &node2
	node1.Right = &node3
	expect := [][]int{
		{1},
		{2, 3},
	}
	actual := levelOrder(&node1)
	assert.Equal(t, expect, actual, "levelOrder execute error")
}
