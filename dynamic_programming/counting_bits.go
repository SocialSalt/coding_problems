package dynamicprogramming

/*
55. Counting Bits

Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := range n + 1 {
		// do a bit shift on i and then check if the
		// bit you deleted was a 1 or not
		// this saves iterating over the binary rep of
		// every number and uses the fact that you
		// need to do the computation for each number
		// anyway
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}
