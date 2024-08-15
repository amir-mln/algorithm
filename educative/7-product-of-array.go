package educative

// TODO: Two Pointers...

var ProductOfArray = map[string]func([]int) []int{
	"two-pointers": func(numbers []int) []int {
		lp, rp := 0, len(numbers)-1
		lSum, rSum := 1, 1
		res := make([]int, len(numbers))

		for i := 0; i < len(res); i++ {
			res[i] = 1
		}

		for lp < len(numbers) && rp >= 0 {
			res[lp] *= lSum
			res[rp] *= rSum
			lSum *= numbers[lp]
			rSum *= numbers[rp]
			lp++
			rp--
		}

		return res
	},
	"cache-slice": func(numbers []int) []int {
		pre := make([]int, len(numbers))
		post := make([]int, len(numbers))

		for i := 0; i < len(pre); i++ {
			if i == 0 {
				pre[i] = 1
			} else {
				pre[i] = numbers[i-1] * pre[i-1]
			}
		}

		for i := len(post) - 1; i >= 0; i-- {
			if i == len(post)-1 {
				post[i] = 1
			} else {
				post[i] = numbers[i+1] * post[i+1]
			}
		}

		res := make([]int, len(numbers))
		for i := 0; i < len(res); i++ {
			res[i] = pre[i] * post[i]
		}

		return res
	},
}
