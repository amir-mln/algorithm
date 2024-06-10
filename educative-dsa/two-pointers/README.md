# Two Pointers

As the name suggests, the two pointers pattern uses two pointers to iterate over an array or list until the conditions of the problem are satisfied. This is useful because it allows us to keep track of the values of two different indexes in a single iteration. Whenever there’s a requirement to find two data elements in an array that satisfy a certain condition, the two pointers pattern should be the first strategy to come to mind.

The pointers can be used to iterate the data structure in one or both directions, depending on the problem statement. For example, to identify whether a string is a palindrome, we can use one pointer to iterate the string from the beginning and the other to iterate it from the end. At each step, we can compare the values of the two pointers and see if they meet the palindrome properties.

The naive approach to solving this problem would be using nested loops, with a time complexity of O(n<sup>2</sup>). However, by using two pointers moving towards the middle from either end, we exploit the symmetry property of palindromic strings. This allows us to compare the elements in a single loop, making the algorithm more efficient with a time complexity ofO(n).

Here’s how the pointers will move along the string:

```
// first iteration
[a, b, c, c, b, a]
p1: a, p2: a, p1 == p2

// second iteration
[a, b, c, c, b, a]
p1: b, p2: b, p1 == p2

// third iteration
[a, b, c, c, b, a]
p1: c, p2: c, p1 == p2
```

## Problems

### 1. Add Binary

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
