package arrays

import (
	"slices"
)

func minimumEffort(tasks [][]int) int {

	slices.SortFunc(tasks, func(a []int, b []int) int {
		firstDiff := a[1] - a[0]
		secondDiff := b[1] - b[0]
		return firstDiff - secondDiff
	})

	energy := 0
	for i := range tasks {
		energy = max(energy+tasks[i][0], tasks[i][1])
	}

	return energy
}
