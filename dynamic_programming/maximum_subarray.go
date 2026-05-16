package dynamicprogramming

import "math"

func maxSubArray(nums []int) int {

	bestSumArray := make([]int, len(nums)+1)
	bestSum := math.MinInt

	// the best sum array is kinda like a prefix array
	// the value bestSumArray[i+1] is the sum of the best
	// subarray up to and including i
	// THIS IS A PREFIX ARRAY PROBLEM BUT!!! we are choosey about
	// what the prefix is
	// the reason we set bestSumArray[i+1] = max(bestSumArray[i]+nums[i], nums[i])
	// is because we want to know if i is part of a generational run
	// or if it's just better to start over at i
	// in the case we don't even know where the best subarray starts
	// but we do know where it ends
	for i := range nums {
		bestSumArray[i+1] = max(bestSumArray[i]+nums[i], nums[i])
		if bestSumArray[i+1] > bestSum {
			bestSum = bestSumArray[i+1]
		}
	}

	return bestSum
}
