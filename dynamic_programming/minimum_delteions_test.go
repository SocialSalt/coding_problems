package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimumDeletions(t *testing.T) {
	var res int
	res = minimumDeletions("aababbab")
	require.Equal(t, 2, res)

	res = minimumDeletions("bbaaaaabb")
	require.Equal(t, 2, res)

}

func TestMinimumDeletionsSimple(t *testing.T) {
	var res int
	res = minimumDeletionsSimple("aababbab")
	require.Equal(t, 2, res)

	res = minimumDeletionsSimple("bbaaaaabb")
	require.Equal(t, 2, res)

}
