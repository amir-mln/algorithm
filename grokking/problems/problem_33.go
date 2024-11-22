package problems

import (
	"math"
	"strconv"
)

/*
# Fraction to Recurring Decimal

Given the two integer values of a fraction, [numerator] and [denominator],
implement a function that returns the fraction in string format. If the
fractional part repeats, enclose the repeating part in parentheses.

Constraints:

 1. denominator != 0
 2. −10**5 ≤ numerator, denominator ≤ 10**5-1
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
