package problems

import "math"

/*
# Repeated DNA Sequences

Given a string, that represents a DNA sequence, and a number k,
return all the contiguous subsequences (substrings) of length k
that occur more than once in the string. The order of the returned
subsequences does not matter. If no repeated substring is found,
the function should return an empty set.

The DNA sequence is composed of a series of nucleotides abbreviated
as  A, C, G, and T. For example, ACGAATTCCG is a DNA sequence. When
studying DNA, it is useful to identify repeated sequences in it.

Constraints:

 1. 1 ≤ len(DNA) ≤ 10000
 2. DNA[i] is either A, C, G, and T.
 3. 1 ≤ k ≤ 10

Description:

We traverse the string by using a sliding window of length [k],
which slides one character forward on each iteration. On each
iteration, we compute the hash of the current k-length substring
in the window. We check if the hash is already present in the
set; If it is, the substring is repeated, so we add it to the
output. Otherwise, the substring has not yet been repeated, so
we add the computed hash to the set. We repeat the above process
for all k-length substrings by sliding the window one character
forward on each iteration. After all k-length substrings have
been evaluated, we return the output.

We need a hash function that helps us achieve constant-time
hashing.For this purpose, we use the polynomial rolling hash
technique:

	H= (c1 * (a ^ k-1)) + (c2 * (a ^ k-2)) + ... + (ck * (a ^ 0))

Here, [a] is a constant, c1, …, ck are the characters in a sequence,
and k is the substring length. Since we only have 4 possible nucleotides,
our a would be 4. We also assign numeric values to the nucleotides,
as shown in the below:

	A->1, C->2, G->3, T->4

We are basically converting the k-length substring window to a
base-4 number of length [k]. The polynomial hash value for the
sequence ATG will be

	H(ATG)=(1×4^2)+(4×4^1)+(3×4^0)=35

To compute the next hash value, we just need to remove the contribution
of the left most character and add the contribution of the rightmost
character. As we remove the left most character, we need to shift the
remaining bases to the left by one position so that the hash corresponds
to the new sliding window. We do this by multiplying the current hash
value by the base value 4.
*/
type Problem40 struct {
	DNA string
	K   int
}

func (p Problem40) Solve() []string {
	const base = 4
	hash := 0
	mp := map[rune]int{
		'A': 1,
		'C': 2,
		'G': 3,
		'T': 4,
	}
	reps := map[int]int{}
	res := make([]string, 0)

	for i := 0; i < len(p.DNA); {
		if i == 0 {
			for j := 0; j < p.K; j++ {
				ch := rune(p.DNA[j])
				expo := float64(p.K - (j + 1))
				hash += mp[ch] * int(math.Pow(base, expo))
			}
			reps[hash]++
			i = p.K
			continue
		}

		leftMostCh := rune(p.DNA[i-p.K])
		rightMostCh := rune(p.DNA[i])
		hash -= mp[leftMostCh] * int(math.Pow(base, float64(p.K-1)))
		hash *= base
		hash += mp[rightMostCh]

		if val, ok := reps[hash]; ok && val == 1 {
			reps[hash]++

			res = append(res, p.DNA[i-p.K+1:i+1])
		} else {
			reps[hash]++
		}
		i++
	}

	return res
}
