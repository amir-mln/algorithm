# Container with Most Water

You’re given an integer array `heights` of length `n`, and there are `n` vertical lines drawn such that the two endpoints of the i<sub>th</sub> line are (`i`, `0`) and (`i`, `height[i]`).

Find two lines from the input array that, together with the x-axis, form a container that holds as much water as possible. Return the maximum amount of water a container can store.

**Constraints:**

- 2 ≤ `n` ≤ 10<sup>3</sup>
- 0 ≤ `heights[i]` ≤ 10<sup>3</sup>

**Hints & Notes:**

1. You can solve this problem using nested for loops with O(n<sup>2</sup>) time complexity.
2. The better approach is to use two pointers; one pointing at the start and the other at the end
3. With the two pointers approach, at the end of each iteration, the pointer that points at the shorter height must be moved. That's because the distance between the two vertical line is multiplied by the shorter height.
