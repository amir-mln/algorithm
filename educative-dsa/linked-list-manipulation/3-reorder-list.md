# Reorder List

Given the head of a singly linked list, reorder the list as if it were folded on itself. For example, if the list is represented as follows:

𝐿<sub>0</sub> → 𝐿<sub>1</sub> → 𝐿<sub>2</sub> → … → 𝐿<sub>n-2</sub> → 𝐿<sub>n-1</sub> → 𝐿<sub>n</sub>

After reordering, it becomes:

𝐿<sub>0</sub> → 𝐿<sub>n</sub> → 𝐿<sub>1</sub> → 𝐿<sub>n-1</sub> → 𝐿<sub>2</sub> → 𝐿<sub>n-2</sub> → …

**Constraints**:

- The range of number of nodes in the list is [1,500].
- −5000 ≤ Node.value ≤ 5000
