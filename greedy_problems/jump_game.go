package greedyproblems

func canJump(nums []int) bool {
	jumps := make([]int, 0, len(nums))
	index := 0

	for index < len(nums)-1 {
		if nums[index] == 0 {
			if len(jumps) > 0 {
				index = index - jumps[len(jumps)-1]
				jumps = jumps[:len(jumps)-1]
			} else {
				return false
			}

		} else {
			jumps = append(jumps, nums[index])
			nums[index] = nums[index] - 1
			index += jumps[len(jumps)-1]
		}
	}
	return true

}

func jump(nums []int) int {
	jumps := 0
	farthest := 0
	end := 0
	for i := range nums {
		farthest = max(farthest, i+nums[i])
		if end >= len(nums)-1 {
			break
		}
		if i == end {
			jumps++
			end = farthest
		}
	}
	return jumps
}
