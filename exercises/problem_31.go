package exercises

/*
# Isomorphic Strings

Given two strings, str1 and str2, check whether two strings are
isomorphic to each other or not.
Two strings are isomorphic if a fixed mapping exists from the
characters of one string to the characters of the other string.
For example, if there are two instances of the character "a" in
the first string, both these instances should be converted to
another character (which could also remain the same character if
"a" is mapped to itself) in the second string. This converted
character should remain the same in both positions of the second
string since there is a fixed mapping from the character "a" in
the first string to the converted character in the second string.
Two different characters cannot map to the same character.
Furthermore, all the instances of a character must be replaced
with another character while protecting the order of characters.

Constraints:

 1. Both the strings consist of valid ASCII characters.
 2. 0 ≤ len(str1), len(str2) ≤ 5000
 3. len(str1) == len(str2)
*/
var Problem31 = map[string]func(srt1, srt2 string) bool{
	"linear-space": func(srt1, srt2 string) bool {
		dict1 := make(map[rune]rune)
		dict2 := make(map[rune]rune)

		for i, r1 := range srt1 {
			r2 := rune(srt2[i])
			m1, ok1 := dict1[r1]
			m2, ok2 := dict2[r2]

			if !ok1 && !ok2 {
				dict1[r1] = r2
				dict2[r2] = r1
				continue
			}

			if ok1 && m1 != r2 || ok2 && m2 != r1 {
				return false
			}
		}

		return true
	},
}
