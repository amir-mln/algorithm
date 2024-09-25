package exercises

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
var Problem32 = map[string]func(srt string) int{
	"linear-space": func(srt string) int {
		dict := make(map[rune]int)
		sum := 0

		for _, r := range srt {
			prev := dict[r]
			dict[r] = prev + 1
		}

		prevOdd := 0
		for _, v := range dict {
			if v%2 != 0 {
				/*
					If we have multiple odd numbers we only want to add
					the maximum odd number. So if the current value is
					less than or equal to the previously added odd number,
					we simply increase sum by the current value minus one.
					this way we can still use the characters that have odd
					occurrences but are less than the maximum odd number.
					And if the current value is more than the previous
					odd number, we simply decrease sum by one, based on the
					same concept.
				*/
				if v <= prevOdd {
					sum += v - 1
					continue
				}

				if prevOdd != 0 {
					sum -= 1
				}

				prevOdd = v
			}

			sum += v
		}

		return sum
	},
}
