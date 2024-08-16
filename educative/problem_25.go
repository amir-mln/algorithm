package educative

import "math"

/*
Largest Rectangle in Histogram

Given a slice of integers, heights, that represents the heights of bars in a histogram,
find the area of the largest rectangle in the histogram, where the rectangle has a constant
width of 1 unit for each bar.
Constraints:

	a) 1 ≤ len(heights) ≤ 10**3
	b) 0 ≤ heights[i] ≤ 10**4
*/
var P25Solutions = map[string]func([]int) int{
	"nested-loop": func(heights []int) int {
		area := 0

		for i := 0; i < len(heights); i++ {
			minHeight := math.MaxInt64
			for j := i; j >= 0 && heights[j] != 0; j-- {
				width := i - j + 1
				minHeight = min(minHeight, heights[j])
				area = max(area, width*minHeight)
			}
		}

		return area
	},
	// "stacks-solution": func(nums []int) int {
	// 	return 0
	// },
}