package findduplicates

func findDuplicateWithMissing(nums []int) int {
	n := len(nums)
	var sum1 int = 0
	var sum2 int = 0

	for _, num := range nums {
		sum1 += num
		sum2 += (num * num)
	}
	e1 := (n * (n + 1)) / 2
	e2 := (n * (n + 1) * (2*n + 1)) / 6
	x := sum1 - e1
	y := sum2 - e2
	result := (y - x*x) / (2 * x)
	return result + x
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}
