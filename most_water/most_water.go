package mostwater

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

func maxArea(height []int) int {
	bestArea := 0
	lower := 0
	upper := len(height) - 1

	for lower < upper {

		h := min(height[lower], height[upper])
		area := h * (upper - lower)
		if area > bestArea {
			bestArea = area
		}

		if height[lower] < height[upper] {
			lower++
		} else {
			upper--
		}

	}

	return bestArea
}
