package exercises

// TODO: Two Pointers...

var SortColors = map[string]func(colors []int){
	"two-pointers-solution": func(colors []int) {
		p1, p2 := 0, 0

		for i := 0; i < 2; i++ {
			p2 = p1 + 1
			for p2 < len(colors) && p1 < len(colors) {
				// p1 should point to a color > i
				if colors[p1] <= i {
					p1++
					continue
				}

				if p2 <= p1 {
					p2 = p1 + 1
					continue
				}

				if colors[p2] != i {
					p2++
					continue
				}

				colors[p1], colors[p2] = colors[p2], colors[p1]
			}
		}
	},
	"three-pointers-solution": func(colors []int) {
		start := 0
		current := 0
		end := len(colors) - 1

		for current <= end {
			if colors[current] == 0 {

				if colors[start] != 0 {
					colors[start], colors[current] = colors[current], colors[start]
				}

				current++
				start++
			} else if colors[current] == 1 {

				current++
			} else {
				if colors[end] != 2 {
					colors[current], colors[end] = colors[end], colors[current]
				}

				end--
			}
		}
	},
}
