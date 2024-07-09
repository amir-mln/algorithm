package binarysearch

import "educative-dsa/common"

var SearchInRotatedSortedArray = map[string]func(nums []int, target int) int{
	"my-binary-search": func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		mid := l + (r-l)/2

		if nums[mid] < nums[l] || nums[mid] > nums[r] {
			for l != r {
				mid = l + (r-l)/2

				if nums[mid] > nums[r] {
					l = mid + 1
				} else {
					r = mid
				}
			}
		}

		if nums[l] <= target && target <= nums[len(nums)-1] {
			r = len(nums) - 1
		}

		if nums[0] <= target && l != 0 && target <= nums[l-1] {
			r = l - 1
			l = 0
		}

		if idx, ok := common.BinarySearchIdx(nums, target, l, r); ok {
			return idx
		} else {
			return -1
		}
	},

	"educative-binary-search": func(nums []int, target int) int {
		start := 0
		end := len(nums) - 1

		for start <= end {
			mid := start + (end-start)/2
			if nums[mid] == target {
				return mid
			}
			if nums[start] <= nums[mid] {
				if nums[start] <= target && target < nums[mid] {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		return -1
	},
}
