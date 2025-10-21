package arraysum

import (
	"math"
	"slices"
)

/*
15.
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	sums := [][]int{}
	for i := 0; i < len(nums)-2; {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s < 0 {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			} else if s > 0 {
				k--
				for k < j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				sums = append(sums, []int{nums[i], nums[j], nums[k]})
				j++
				for j < len(nums)-1 && nums[j-1] == nums[j] {
					j++
				}
			}
		}
		i++
		for i < len(nums)-2 && nums[i-1] == nums[i] {
			i++
		}
	}
	return sums
}

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	closest := math.MaxInt
	for i := 0; i < len(nums)-2; {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(s-target)) < math.Abs(float64(closest-target)) {
				closest = s
			}
			if (s - target) < 0 {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			} else if (s - target) > 0 {
				k--
				for k < j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				return target
			}
		}
		i++
		for i < len(nums)-2 && nums[i-1] == nums[i] {
			i++
		}
	}
	return closest
}

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	sums := [][]int{}

	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			k := j + 1
			l := len(nums) - 1
			for k < l {
				s := nums[i] + nums[j] + nums[k] + nums[l]
				if s < target {
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
				} else if s > target {
					l--
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else {
					sums = append(sums, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
				}
			}
			j++
			for j < len(nums)-2 && nums[j-1] == nums[j] {
				j++
			}
		}
		i++
		for i < len(nums)-3 && nums[i-1] == nums[i] {
			i++
		}
	}
	return sums
}
