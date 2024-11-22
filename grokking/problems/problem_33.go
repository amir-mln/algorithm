package problems

import (
	"math"
	"strconv"
)

/*
# Longest Palindrome

Given a string s that only contains alphabets, return the
length of the longest palindrome that may be composed using
those letters. Letters are case-sensitive. Hence, combinations
such as “Aa” are not considered palindromes. Also, it's ok
to not use every instance of a character in the palindrome.
For example, if you have three d's, you can only use 2 d's
in you calculations.

Constraints:

 1. str consist of valid ASCII characters.
 2. 0 ≤ len(str) ≤ 1000
*/
type Problem33 struct {
	Numerator   int
	Denominator int
}

func (p *Problem33) Solve() string {
	result := ""
	remainders := make(map[int]int)

	if p.Numerator == 0 {
		return "0"
	}

	quotient := p.Numerator / p.Denominator
	if p.Numerator < 0 || p.Denominator < 0 {
		if quotient <= 0 {
			result += "-"
			quotient = int(math.Abs(float64(quotient)))
		}
		p.Numerator = int(math.Abs(float64(p.Numerator)))
		p.Denominator = int(math.Abs(float64(p.Denominator)))
	}

	remainder := p.Numerator % p.Denominator * 10
	result += strconv.Itoa(quotient)
	if remainder == 0 {
		return result
	}

	result += "."
	for remainder != 0 {
		if beginning, ok := remainders[remainder]; ok {
			left := result[0:beginning]
			right := result[beginning:]
			result = left + "(" + right + ")"
			return result
		}

		remainders[remainder] = len(result)
		quotient = remainder / p.Denominator
		result += strconv.Itoa(quotient)
		remainder = (remainder % p.Denominator) * 10
	}

	return result
}
