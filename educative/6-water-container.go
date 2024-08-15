package educative

// TODO: Two Pointers...

var ContainerWithMostWater = map[string]func(heights []int) int{
	"nested-iteration": func(heights []int) int {
		area := 0

		for i := 0; i < len(heights); i++ {
			for j := i + 1; j < len(heights); j++ {
				area = max(area, (j-i)*min(heights[j], heights[i]))
			}
		}

		return area
	},
	"two-pointers": func(heights []int) int {
		p1, p2, area := 0, len(heights)-1, 0

		for p1 < p2 {
			area = max(area, (p2-p1)*min(heights[p1], heights[p2]))

			if heights[p2] <= heights[p1] {
				p2--
			} else {
				p1++
			}
		}

		return area
	},
}
