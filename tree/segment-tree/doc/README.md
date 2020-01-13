## Segment Tree | Set 1 (Sum of given range)

### Time Complexity
Time Complexity for tree construction is $O(n)$. There are total $2n-1$ nodes, and value of every node is calculated only once in tree construction.

Time complexity to query is $O(Logn)$. To query a sum, we process at most four nodes at every level and number of levels is $O(Logn)$.

The time complexity of update is also $O(Logn)$. To update a leaf value, we process one node at every level and number of levels is $O(Logn)$.