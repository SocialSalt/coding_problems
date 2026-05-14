package greedyproblems

func minMoves(nums []int, limit int) int {
	l := len(nums)
	diffs := make([]int, 2*limit+2)

	for i := range l / 2 {
		a := nums[i]
		b := nums[l-i-1]

		// We're going to use a difference array to find the optimal number of changes
		// required to get this array matching. So we need to set up the difference
		// array correctly. We want to set up the array so that sum(array[0:c+1]) gives
		// the total number of changes needed if c is the best sum to choose. For each
		// pair we're precomputing the number of changes required for each range.
		//
		// for the pair (a,b) we have the following outcomes
		// entering the range [2,min(a,b)+1) we'll need 2 chanages
		diffs[2] += 2
		// in the range [min(a,b)+1, a+b) we'll only need one change
		diffs[min(a, b)+1] -= 1
		// at (a+b) we need no changes
		diffs[a+b] -= 1
		// in the range (a+b, max(a,b) + limit] we'll need one change
		diffs[a+b+1] += 1
		// in the range (max(a,b) + limit, 2*limit] we'll need two changes again
		diffs[max(a, b)+limit+1] += 1
	}

	bestCount := l + 1
	runningCount := 0
	for i := 2; i < 2*limit+1; i++ {
		runningCount += diffs[i]
		bestCount = min(bestCount, runningCount)
	}

	return bestCount
}
