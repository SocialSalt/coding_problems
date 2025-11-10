package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRob(t *testing.T) {
	var res int
	var root TreeNode
	root = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 1},
		},
	}
	res = rob(&root)
	require.Equal(t, 7, res)

	root = TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		},
	}
	res = rob(&root)
	require.Equal(t, 7, res)
}
