# Circular Array Loop

We are given a circular array of non-zero integers, `nums`, where each integer represents the number of steps to be taken either forward or backward from its current index. Positive values indicate forward movement, while negative values imply backward movement. When reaching either end of the array, the traversal wraps around to the opposite end.

The input array may contain a cycle, which is a sequence of indexes characterized by the following:

- The sequence starts and ends at the same index.
- The length of the sequence is at least two.
- The loop must be in a single direction, forward or backward.

Your task is to determine if nums has a cycle. Return `true` if there is a cycle. Otherwise return `false`.

**Constraints**:

- 1 ≤ `len(nums)` ≤ 10<sup>5</sup>
- -5000 ≤ `nums[i]` ≤ 5000
- `nums[i]` ≠ 0
- A cycle in the array may begin from any index in the array.
