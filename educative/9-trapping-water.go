package educative

// TODO: Two Pointers...

var TrappingRainWater = map[string]func(heights []int) int{
	"two-pointers": func(heights []int) int {
		sum := 0
		tempSum := 0
		i := 0
		lastPeak := -1

	beginning:
		// i will point at the first peak after this
		for i < len(heights)-1 && heights[i] <= heights[i+1] {
			i++
		}

		for j := i + 1; j < len(heights); j++ {
			if heights[j] >= heights[i] {
				i = j
				sum += tempSum
				tempSum = 0
				lastPeak = -1
				goto beginning
			}

			if heights[j] > heights[j-1] && (lastPeak == -1 || heights[lastPeak] < heights[j]) {
				lastPeak = j
			}

			tempSum += heights[i] - heights[j]
		}

		if tempSum > 0 && lastPeak != -1 {
			extra := 0

			for k := i + 1; k < len(heights); k++ {
				if k >= lastPeak || heights[k] >= heights[lastPeak] {
					extra += heights[i] - heights[k]
				} else {
					extra += heights[i] - heights[lastPeak]
				}
			}

			tempSum -= extra
			sum += tempSum

			if lastPeak < len(heights) {
				i = lastPeak
				lastPeak = -1
				tempSum = 0
				goto beginning
			}
		}

		return sum
	},
}
