package segmenttree

// Node a segment-tree's node
type Node struct {
	Left  *Node
	Right *Node
	Start int
	End   int
	Value int
}

// SegmentTree with Root node
type SegmentTree struct {
	Root *Node
}

// New Function to construct segment tree from given array.
func New(arr []int) *SegmentTree {
	root := &Node{}
	st := &SegmentTree{
		root,
	}
	st.build(arr, 0, len(arr)-1, root)
	return st
}

// UpdateValue The function to update a value in input array and segment tree.
// It uses update() to update the value in segment tree.
func (st *SegmentTree) UpdateValue(start int, end int, diff int) {
	st.update(start, end, diff, st.Root)
}

// GetSum Return sum of elements in range from index start (quey start) to end (query end).
// It use query() to query the sum in segment tree.
func (st *SegmentTree) GetSum(start int, end int) int {
	return st.query(start, end, st.Root)
}

// build Function to construct segment tree from given array.
func (st *SegmentTree) build(arr []int, start int, end int, node *Node) int {
	node.Start = start
	node.End = end
	// If there is one element in array, store it in current node of segment tree and return
	if start == end {
		node.Value = arr[start]
		return node.Value
	}
	// If there are more than one elements, then recur for left and
	// right subtrees and store the sum of values in this node
	mid := getMid(start, end)
	node.Value = st.build(arr, start, mid, node.buildLeft()) + st.build(arr, mid+1, end, node.buildRight())
	return node.Value
}

// update A recursive function to update the nodes from index start to end.
func (st *SegmentTree) update(start int, end int, diff int, node *Node) {
	// If the input index lies outside the range of this segment, return
	if node.Start > end || node.End < start {
		return
	}
	// calculate which interval should be updated in current node
	if node.Start > start {
		start = node.Start
	}
	if node.End < end {
		end = node.End
	}
	// If the input index is in range of this node, then update the value of the node and its children
	node.Value += (end - start + 1) * diff
	if node.Start != node.End {
		st.update(start, end, diff, node.Left)
		st.update(start, end, diff, node.Right)
	}
}

// query A recursive function to get the sum of values in the given range of the array.
func (st *SegmentTree) query(start int, end int, node *Node) int {
	// If segment of this node is a part of given range, then return the sum of the segment
	if start <= node.Start && node.End <= end {
		return node.Value
	}
	// If segment of this node is outside the given range, return
	if node.Start > end || node.End < start {
		return 0
	}
	// If a part of this segment overlaps with the given range
	return st.query(start, end, node.Left) + st.query(start, end, node.Right)
}

func (node *Node) buildLeft() *Node {
	if node.Left == nil {
		node.Left = &Node{}
	}
	return node.Left
}

func (node *Node) buildRight() *Node {
	if node.Right == nil {
		node.Right = &Node{}
	}
	return node.Right
}

func getMid(start int, end int) int {
	return (start + end) / 2
}
