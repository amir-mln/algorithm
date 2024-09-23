package exercises

import "math"

// TODO: Fast & Slow Pointers

var CircularArrayLoop = map[string]func(nums []int) bool{
	"fast-slow-pointer": func(nums []int) bool {
		getNextIndex := func(idx int) int {
			newIdx := nums[idx] + idx

			for newIdx < 0 {
				newIdx += len(nums)
			}

			for newIdx >= len(nums) {
				newIdx -= len(nums)
			}

			return newIdx
		}

		for i, v := range nums {
			var dir int

			if v > 0 {
				dir = 1
			} else {
				dir = -1
			}

			sp, fp := i, i

			for {

				sp = getNextIndex(sp)

				before := fp
				fp = getNextIndex(fp)
				if math.Signbit(float64(nums[fp])) != math.Signbit(float64(dir)) || fp == before {
					break
				}

				before = fp
				fp = getNextIndex(fp)
				if math.Signbit(float64(nums[fp])) != math.Signbit(float64(dir)) || fp == before {
					break
				}

				if fp == sp {
					return true
				}
			}
		}

		return false
	},
}
