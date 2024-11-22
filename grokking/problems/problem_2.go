package problems

import "math"

/*
String to Integer
#Two_Pointers

Write a function that converts a string to a 32–bit signed integer. It is equivalent to the
atoi function. Any leading whitespace must be ignored. The function should also check for +
or - sign. There is no need to add a '+' sign in front of the resulting integer. For example,
"-25" converts to -25, and "+91"converts to 91. However, if neither is present, assume the
result is positive. If the first non-space character is neither a sign character, nor a non-
digit character, the function stops processing further and returns 0. For example, the strings
"One1" and " .5" convert to 0. Additionally, if the first non-space character is a sign character,
'+' or '-', and the next character is a non-digit character (including the space character),
the function returns0. For example, the strings +.23", " - 49", and " +R345" convert to 0. Then,
the function should read digit to the end of string or when a non-digit character is encountered.
Leading zeros from the digits must ber ignored as well. If the end result goes out of the range of
a 32–bit signed integer, return maximum signed 32 bit integers for positive numbers and minimum signed
32 bit integer for negative numbers.

Constraints:

 1. 0 ≤ len(s) ≤ 200
 2. The string s may have Digit characters from 0-9.
 3. The string s may have Non-digit characters, including lower-case and upper-case English letters,
    space character ' ', period ., and sign characters + and -.
*/
var Problem2 = map[string]func(s string) int32{
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
