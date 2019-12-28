### Selection Sort
- **Ref**: https://www.geeksforgeeks.org/selection-sort/
- **Time Complexity**: $O(n^2)$ as there are two nested loops.
- **Auxiliary Space**: $O(1)$
The good thing about selection sort is it never makes more than O(n) swaps and can be useful when memory write is a costly operation.
- **Sorting In Place**: Yes
- **Stable**: No

### Stable Selection Sort
- **Ref**: https://www.geeksforgeeks.org/stable-selection-sort/

### Bubble Sort
- **Ref**: https://www.geeksforgeeks.org/bubble-sort/
- **Worst and Average Case Time Complexity**: $O(n^2)$. 
  Worst case occurs when array is reverse sorted.
- **Best Case Time Complexity**: $O(n)$. 
  Best case occurs when array is already sorted.
- **Auxiliary Space**: $O(1)$
- **Boundary Cases**: Bubble sort takes minimum time (Order of $n$) when elements are already sorted.
- **Sorting In Place**: Yes
- **Stable**: Yes

### Insertion Sort
- **Ref**: https://www.geeksforgeeks.org/insertion-sort/
- **Time Complexity**: $O(n^2)$
- **Auxiliary Space**: $O(1)$
- **Boundary Cases**: Insertion sort takes maximum time to sort if elements are sorted in reverse order. And it takes minimum time (Order of $n$) when elements are already sorted.
- **Algorithmic Paradigm**: Incremental Approach
- **Sorting In Place**: Yes
- **Stable**: Yes
- **Online**: Yes
- **Uses**: Insertion sort is used when number of elements is small. It can also be useful when input array is almost sorted, only few elements are misplaced in complete big array.

### Merge Sort
- **Ref**: https://www.geeksforgeeks.org/merge-sort/
- **Time Complexity**: Sorting arrays on different machines. Merge Sort is a recursive algorithm and time complexity can be expressed as following recurrence relation.
$$
T(n) = 2T(n/2) + \Theta(n)
$$
The above recurrence can be solved either using Recurrence Tree method or Master method. It falls in case II of Master Method and solution of the recurrence is $\Theta(n\log n)$.
Time complexity of Merge Sort is $\Theta(n\log n)$ in all 3 cases (worst, average and best) as merge sort always divides the array into two halves and take linear time to merge two halves.
- **Auxiliary Space**: $O(n)$
- **Algorithmic Paradigm**: Divide and Conquer
- **Sorting In Place**: No in a typical implementation
- **Stable**: Yes
- **Applications of Merge Sort**
  - [Merge Sort is useful for sorting linked lists in $O(n\log n)$ time](https://www.geeksforgeeks.org/merge-sort-for-linked-list/)
  - [Inversion Count Problem](https://www.geeksforgeeks.org/counting-inversions/)
  - [External Sorting](http://en.wikipedia.org/wiki/External_sorting)

### Quick Sort
- **Ref**: https://www.geeksforgeeks.org/quick-sort/
- **Analysis of QuickSort**: 
  Time taken by QuickSort in general can be written as following.
  $$
  T(n) = T(k) + T(n-k-1) + \theta(n)
  $$
  The first two terms are for two recursive calls, the last term is for the partition process. $k$ is the number of elements which are smaller than pivot.
  The time taken by QuickSort depends upon the input array and partition strategy. Following are three cases.
- **Worst Case**: 
  The worst case occurs when the partition process always picks greatest or smallest element as pivot. If we consider above partition strategy where last element is always picked as pivot, the worst case would occur when the array is already sorted in increasing or decreasing order. Following is recurrence for worst case.
  $$
  T(n) = T(0) + T(n-1) + \theta(n)
  $$
  which is equivalent to 
  $$
  T(n) = T(n-1) + \theta(n)
  $$
  The solution of above recurrence is $\theta(n^2)$.
- **Best Case**: 
  The best case occurs when the partition process always picks the middle element as pivot. Following is recurrence for best case.
  $$
  T(n) = 2T(n/2) + \theta(n)
  $$
  The solution of above recurrence is $\theta(n\log n)$. It can be solved using case 2 of Master Theorem.
- **Average Case**:
  To do average case analysis, we need to consider all possible permutation of array and calculate time taken by every permutation which doesn’t look easy.
  We can get an idea of average case by considering the case when partition puts $O(n/9)$ elements in one set and $O(9n/10)$ elements in other set. Following is recurrence for this case.
  $$
  T(n) = T(n/9) + T(9n/10) + \theta(n)
  $$
  Solution of above recurrence is also $O(n\log n)$
  Although the worst case time complexity of QuickSort is $O(n^2)$ which is more than many other sorting algorithms like Merge Sort and Heap Sort, QuickSort is faster in practice, because its inner loop can be efficiently implemented on most architectures, and in most real-world data. QuickSort can be implemented in different ways by changing the choice of pivot, so that the worst case rarely occurs for a given type of data. However, merge sort is generally considered better when data is huge and stored in external storage.
- **Stable**: No
- **Sorting In Place**: Yes
