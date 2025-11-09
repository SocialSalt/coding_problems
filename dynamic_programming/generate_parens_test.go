package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateParenthesis(t *testing.T) {
	var res []string
	res = generateParenthesis(3)
	require.Equal(
		t,
		[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		res,
	)
}
