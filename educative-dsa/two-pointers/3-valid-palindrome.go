package twopointers

var ValidPalindrome = map[string]func(str1 string) bool{
	"linear-solution": func(str1 string) bool {
		runes := []rune(str1)
		l, r := 0, len(str1)-1

		for l <= r {
			if runes[l] != runes[r] {
				return false
			}

			l++
			r--
		}

		return true
	},
}
