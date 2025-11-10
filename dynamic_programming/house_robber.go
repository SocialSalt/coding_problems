package dynamicprogramming

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }

337. House Robber III
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.

Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night.

Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	// Approach:
	// For any root passed into maxSum there are two possible values to return:
	// either the root is included in the sum or it is excluded from the sum. If
	// it is included then we can only used the excluded values for the children.
	// If it is excluded then we can use either the included or excluded values for
	// the children, whichever is larger

	var maxSum func(*TreeNode) (int, int)
	maxSum = func(r *TreeNode) (int, int) {
		if r == nil {
			return 0, 0
		}
		leftIncluded, leftExcluded := maxSum(r.Left)
		rightIncluded, rightExcluded := maxSum(r.Right)
		selfIncluded := r.Val + leftExcluded + rightExcluded
		selfExcluded := max(leftIncluded, leftExcluded) + max(rightIncluded, rightExcluded)
		return selfIncluded, selfExcluded
	}

	rootIncluded, rootExcluded := maxSum(root)
	return max(rootIncluded, rootExcluded)

}
