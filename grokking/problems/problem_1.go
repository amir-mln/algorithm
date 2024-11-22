package problems

import (
	"fmt"
	"strings"
)

/*
Given two binary strings. return their sum as a binary string.

Constraints:

 1. 1 <= len(str1), len(str2) <= 500
 2. Input strings contain 0 and 1 only
 3. Strings can not contain leading zeros except the string representing the binary form of 0

Category: Two Pointers
*/
var Problem1 = map[string]func(str1, str2 string) string{
	"linear-solution": func(str1, str2 string) string {
		p1, p2 := len(str1)-1, len(str2)-1
		carry := 0
		res := make([]string, max(len(str1), len(str2))+1)

		for i := len(res) - 1; i > 0; i-- {
			var d1, d2 int

			if p1 >= 0 {
				d1 = int(str1[p1] - '0')
				p1--
			}

			if p2 >= 0 {
				d2 = int(str2[p2] - '0')
				p2--
			}

			sum := d1 + d2 + carry
			carry = sum / 2
			res[i] = fmt.Sprint(sum % 2)
		}

		if carry != 0 {
			res[0] = fmt.Sprint(carry)
		}

		return strings.Join(res, "")
	},
}
