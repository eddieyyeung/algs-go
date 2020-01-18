# Dynamic Programming

Two patterns of solving DP problem:
- Tabulation: Bottom Up
- Memoization: Top Down

<!-- |   | Tabulation | Memoization  |
|---|---|---|
| State | State transition relation is difficult to think | State transition relation is easy to think |
 -->

Two properties in Dynamic Programming
- [Overlaping Subproblems](https://www.geeksforgeeks.org/overlapping-subproblems-property-in-dynamic-programming-dp-1/)
  Like Divide and Conquer, Dynamic Programming combines solutions to sub-problems. 
- [Optimal Substructure](https://www.geeksforgeeks.org/optimal-substructure-property-in-dynamic-programming-dp-2/)
  A given problems has Optimal Substructure Property if optimal solution of the given problem can be obtained by using optimal solutions of its subproblems.

How to solve a Dynamic Programming Problem ?
1. How to classify a problem as a Dynamic Programming Problem?
2. Deciding the state
   **State** A state can be defined as the set of parameters that can uniquely identify a certain position or standing in the given problem. This set of parameters should be as small as possible to reduce state space.
3. Formulating a relation among the states
4. Adding memoization or tabulation for the state