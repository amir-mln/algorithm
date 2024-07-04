package twopointers

import "math"

var StringAtoi = map[string]func(s string) int32{
	"linear-solution": func(s string) int32 {
		var res int64
		f, l := -1, -1

		digits := map[rune]int64{
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
			'0': 0,
		}

		sign := '+'
		const space = ' '

		for i, r := range s {
			if _, ok := digits[r]; ok {
				if f == -1 {
					f = i
				}

				l = i
				continue
			}

			if r == space && f == -1 && l == -1 {
				continue
			}

			if (r == '+' || r == '-') && f == -1 && l == -1 {
				sign = r
				continue
			}

			break
		}

		if l == -1 && f == -1 {
			return 0
		}

		for i := 0; i <= l-f; i++ {
			res += int64(math.Pow10(l-f-i)) * digits[rune(s[f+i])]
		}

		if sign == '-' {
			res = -res
		}

		if res < math.MinInt32 {
			res = math.MinInt32
		}

		if res > math.MaxInt32 {
			res = math.MaxInt32
		}

		return int32(res)
	},
}
