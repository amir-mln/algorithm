# Valid Palindrome

Given two binary strings `str1` and `str2`. return their sum as a binary string.

**Constraints**:

1. 1 <= `len(str1)`, `len(str2)` <= 500
2. Input strings contain `0` and `1` only
3. Any strings must not contain leading zeros except the string representing the binary form of `0`

**Example**

Inputs: `"100"` , `"111"`

Output: `"1011"`

<details>
<summary><b>Hints & Notes</b></summary>

1. We need to traverse the input strings in reverse order with help of two pointers that point at the end of each input string.
2. We Also need variable to hold the carry value in each iteration.
3. before returning the result, we should check if the carry has a non zero value.

</details>
