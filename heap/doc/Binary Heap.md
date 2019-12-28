## Binary Heap

### Array Representation Of Binary Heap
https://www.geeksforgeeks.org/array-representation-of-binary-heap/
The root element will be at $Arr[0]$.
Below table shows indexes of other nodes for the $i^{th}$ node, i.e., $Arr[i]$:
| Represent | Desc |
|-|-|
| $Arr[(i-1)/2]$ | Returns the parent node |
| $Arr[(2*i)+1]$ | Returns the left child node |
| $Arr[(2*i)+2]$ | Returns the right child node |


