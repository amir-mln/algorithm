package educative

import (
	"slices"
)

// TODO: Two Pointers...

var ThreeSum = map[string]func(nums []int, target int) bool{
	"pointers-with-sorting": func(nums []int, target int) bool {
		slices.Sort(nums)

		for i := 0; i < len(nums)-2; i++ {
			for p1, p2 := i+1, len(nums)-1; p1 < p2; {
				sum := nums[i] + nums[p1] + nums[p2]

				if sum == target {
					return true
				}

				if sum > target {
					p2--
				} else if sum < target {
					p1++
				}
			}
		}

		return false
	},
}
