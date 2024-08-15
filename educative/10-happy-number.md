# Happy Number

Write an algorithm to determine if a number `𝑛` is a happy number. We use the following process to check if a given number is a happy number:

- Starting with the given number `𝑛`, replace the number with the sum of the squares of its digits.
- Repeat the process until:
  - The number equals 1, which will depict that the given number `𝑛` is a happy number.
  - The number enters a cycle, which will depict that the given number `𝑛` is not a happy number.

Return `true` if `𝑛` is a happy number, and FALSE if not.

**Constraints**:

- `𝑛` is a non-zero positive `32-bit` integer
