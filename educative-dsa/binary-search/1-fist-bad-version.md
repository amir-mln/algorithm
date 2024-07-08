# First Bad Version

The latest version of a software product fails the quality check. Since each version is developed upon the previous one, all the versions created after a bad version are also considered bad.

Suppose you have `n` versions with the IDs [1,2,...,n], and you have access to an API function that returns `true` if the argument is the ID of a bad version.

Find the first bad version that is causing all the later ones to be bad. Additionally, the solution should also return the number of API calls made during the process and should minimize the number of API calls too.

**Constraints**:

- 1 ≤ `first bad version` ≤ `n` ≤ 2<sup>31</sup> - 1

1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
1 + 1 + 1 + 1 + 1
