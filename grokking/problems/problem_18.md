# Reverse Nodes In Even Length Groups

Given the head of a linked list, the nodes in it are assigned to each group in a sequential manner. The length of these groups follows the sequence of natural numbers. Natural numbers are positive whole numbers denoted by (1,2,3,4...).

In other words:

- The 1<sub>st</sub> node is assigned to the first group.
- The 2<sub>nd</sub> and 3<sub>rd</sub> nodes are assigned to the second group.
- The 4<sub>th</sub>, 5<sub>th</sub>, and 6<sub>th</sub>nodes are assigned to the third group, and so on.

Your task is to reverse the nodes in each group with an even number of nodes and return the head of the modified linked list.

Note that the length of the last group may be less than or equal to 1 + the length of the second to the last group.

**Constraints**:

- 1 ≤ Number of nodes ≤ 500
- 0 ≤ Node.data ≤ 10<sup>3</sup>
