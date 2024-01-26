package same_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameTree(t *testing.T) {
	node1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node1.Left = node2
	node1.Right = node3
	expect := true
	actual := isSameTree2(node1, node1)
	assert.Equal(t, expect, actual, "isSameTree2 execute error")
}
