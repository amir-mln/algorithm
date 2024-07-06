# String to Integer (atoi)

Write a function that converts a string to a 32–bit signed integer. It is equivalent to the atoi function in C/C++.

Any leading whitespace must be ignored. The function should also check for `+` or `-` sign. There is no need to add a '+' sign in front of the resulting integer. For example, `"-25"` converts to -25, and `"+91"`converts to 91. However, if neither is present, assume the result is positive.

If the first non-space character is neither a sign character, nor a non-digit character, the function stops processing further and returns `0`. For example, the strings `"One1"`and `" .5"`convert to 0. Additionally, if the first non-space character is a sign character, '+' or '-', and the next character is a non-digit character (including the space character), the function returns`0`. For example, the strings `+.23"`, `" - 49"`, and `" +R345"` convert to `0`.

Then, the function should read digit to the end of string or when a non-digit character is encountered. Leading zeros from the digits must ber ignored as well.

If the end result goes out of the range of a 32–bit signed integer, return maximum signed 32 bit integers for positive numbers and minimum signed 32 bit integer for negative numbers.

**Constraints**

- 0 ≤ `len(s)` ≤ 200

- The string s may have:

  1.  Digit characters from 0-9.
  2.  Non-digit characters, including lower-case and upper-case English letters, space character `' '`, period `.`, and sign characters `+` and `-`.

**Examples**

1.  Input: " 1325038"
    Output: "1325038"

2.  Input: " One229988"
    Output: "0"

3.  Input: " One229988"
    Output: "0"
