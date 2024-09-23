package exercises

// TODO: Fast & Slow Pointers

var HappyNumber = map[string]func(n uint32) bool{
	"fast-slow": func(n uint32) bool {
		calcDigitsSqrt := func(n uint32) uint32 {
			sum := uint32(0)

			for n > 0 {
				sum += (n % 10) * (n % 10)
				n = n / 10
			}

			return uint32(sum)
		}

		sp, fp := n, n

		for {
			sp = calcDigitsSqrt(sp)
			fp = calcDigitsSqrt(calcDigitsSqrt(fp))

			if fp == 1 {
				return true
			}

			if fp == sp {
				return false
			}
		}
	},
}
