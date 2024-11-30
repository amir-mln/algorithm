package problems

/*
# Group Anagram

Given a list of words or phrases, group the words that are anagrams
of each other. An anagram is a word or phrase formed from another
word by rearranging its letters.

Constraints:

 1. 1 ≤ len(words) ≤ 1000
 2. 0 ≤ len(word[i]) ≤ 100
 3. Each word consists of lowercase english letters

Description:

The naive approach would be to use the sorted copy of each word as
the key to the hash map. We need to find a way to create the keys
without sorting the words. Since each word consists of only english
letters, we can use a fixed size array to count the frequency of
characters in each word. Then, as arrays are comparable types in go,
we can use them as a key instead of the sorted string.If two words
have the same key they belong to the same group.
*/
type Problem39 []string

func (p Problem39) Solve() [][]string {
	mp := make(map[[26]int]*[]string)

	for _, word := range p {
		freq := [26]int{}

		for _, ch := range word {
			freq[ch-'a']++
		}

		if slc, ok := mp[freq]; ok {
			*slc = append(*slc, word)
		} else {
			mp[freq] = &[]string{word}
		}
	}

	res := make([][]string, 0)
	for _, v := range mp {
		res = append(res, *v)
	}

	return res
}
