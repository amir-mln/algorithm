package exercises

/*
Title:

	Valid Palindrome

Tags:

	#Two_Pointers

Description:

	Write a function that takes a string, `s`, as an input and
	determines whether or not it is a palindrome.

Constraints:

	a) 1 ≤ `len(s)` ≤ 2×10**5
	b) The string `s` will contain English uppercase and lowercase
	letters, digits, and spaces.
*/
var Problem3 = map[string]func(str1 string) bool{
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
