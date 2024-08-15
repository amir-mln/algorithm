# Remove nth Node from End of List

Given a singly linked list, remove the Nth node from the end of the list and return its head.

**Constraints**:

1. The number of nodes in the list is k.
2. 1 ≤ `k` ≤ 10<sup>3</sup>
3. -10<sup>3</sup> ≤ Node.Value ≤ 10<sup>3</sup>
4. 1 ≤ `n` ≤ `k`

**Example**

Input: n= 3, list= [23, 25, 10, 5, 8, 30, 21]

Output: [23, 25, 10, 5, 30, 21]

<details>
<summary><b>Hints & Notes</b></summary>

1. You need two pointers where the second pointer is `n` steps ahead of the first one.
2. Pay attention to the last constraint; `n` is at most equal to number list nodes.
3. You can iterate through the list to get the number of nodes (`k`). Once you have the `k` you can count from the head node to reach to the node that must be removed. This approach does not use two pointers but has the same time and space complexity.
</details>
