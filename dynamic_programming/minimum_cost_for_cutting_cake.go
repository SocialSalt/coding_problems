package dynamicprogramming

import (
	"math"
)

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {

	memo := [21][21][21][21]int{}

	// you had the right idea at first but it turns out that allocating new arrays every time is really slow
	var chooseCut func(int, int, int, int, []int, []int) int
	chooseCut = func(hstart int, hend int, vstart int, vend int, horizontalCut []int, verticalCut []int) int {
		if hend-hstart == 1 && vend-vstart == 1 {
			return 0
		}
		if memo[hstart][hend][vstart][vend] != 0 {
			return memo[hstart][hend][vstart][vend]
		}
		cost := math.MaxInt

		for i := hstart + 1; i < hend; i++ {
			c := horizontalCut[i-1]
			c += chooseCut(hstart, i, vstart, vend, horizontalCut, verticalCut)
			c += chooseCut(i, hend, vstart, vend, horizontalCut, verticalCut)
			cost = min(c, cost)
		}
		for i := vstart + 1; i < vend; i++ {
			c := verticalCut[i-1]
			c += chooseCut(hstart, hend, vstart, i, horizontalCut, verticalCut)
			c += chooseCut(hstart, hend, i, vend, horizontalCut, verticalCut)
			cost = min(c, cost)
		}
		memo[hstart][hend][vstart][vend] = cost
		return cost
	}

	return chooseCut(0, m, 0, n, horizontalCut, verticalCut)
}
