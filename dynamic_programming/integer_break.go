package dynamicprogramming

/*
55. Integer Break

Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.

Return the maximum product you can get.
*/

func integerBreak(n int) int {
	// Approach
	// using a memo here is key, it saves recomputing a lot of products.
	// basically starting with subtracting 2 up to n/2 you check what the product is
	// 2 and 3 are sort of degenrate cases, and hard coding them saves some fenagaling
	// with bounds of the for loop and whatnot.

	memo := make(map[int]int)
	memo[2] = 1
	memo[3] = 2

	var bestProduct func(int) int
	bestProduct = func(v int) int {
		if b, ok := memo[v]; ok {
			return b
		}
		best := 0
		for i := 2; i <= v/2; i++ {
			r := v - i
			bestR := max(r, bestProduct(r))
			bestI := max(i, bestProduct(i))
			best = max(best, bestR*bestI)
		}
		memo[v] = best
		return best
	}
	return bestProduct(n)
}
