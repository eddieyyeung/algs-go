package tree

type SegmentTreeNode struct {
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
	Value int // 当前节点的值
	Add   int // 懒惰标记
}

type SegmentTree struct {
	Head *SegmentTreeNode
	Size int
}

func NewSegmentTree(size int) *SegmentTree {
	return &SegmentTree{
		Head: &SegmentTreeNode{},
		Size: size,
	}
}

// leftNum 和 rightNum 表示左右孩子区间的叶子节点数量
// 因为如果是「加减」更新操作的话，需要用懒惰标记的值✖️叶子节点的数量
func (st *SegmentTree) pushDown(node *SegmentTreeNode, leftNum int, rightNum int) {
	// 动态开点
	if node.Left == nil {
		node.Left = &SegmentTreeNode{}
	}
	if node.Right == nil {
		node.Right = &SegmentTreeNode{}
	}
	// 如果 add 为 0，表示没有标记
	if node.Add == 0 {
		return
	}
	// 注意：当前节点加上标记值✖️该子树所有叶子节点的数量
	node.Left.Value += node.Add * leftNum
	node.Right.Value += node.Add * rightNum
	// 把标记下推给孩子节点
	// 对区间进行「加减」的更新操作，下推懒惰标记时需要累加起来，不能直接覆盖
	node.Left.Add += node.Add
	node.Right.Add += node.Add
	// 取消当前节点标记
	node.Add = 0
}

// 在区间 [start, end] 中更新区间 [l, r] 的值，将区间 [l, r] + val
// 对于上面的例子，应该这样调用该函数：update(root, 0, 4, 2, 4, 1)
func (st *SegmentTree) update(node *SegmentTreeNode, start int, end int, left int, right int, val int) {
	// 找到满足要求的区间
	if left <= start && end <= right {
		// 区间节点加上更新值
		// 注意：需要 * 该子树所有叶子节点
		// TODO
		node.Value += (end - start + 1) * val

		node.Add += val
		return
	}
	mid := (start + end) >> 1
	// 下推标记
	// mid - start + 1：表示左孩子区间叶子节点数量
	// end - mid：表示右孩子区间叶子节点数量
	st.pushDown(node, mid-start+1, end-mid)
	// [start, mid] 和 [l, r] 可能有交集，遍历左孩子区间
	if left <= mid {
		st.update(node.Left, start, mid, left, right, val)
	}
	// [mid + 1, end] 和 [l, r] 可能有交集，遍历右孩子区间
	if right > mid {
		st.update(node.Right, mid+1, end, left, right, val)
	}
	// 向上更新
	st.pushUp(node)
}

func (st *SegmentTree) pushUp(node *SegmentTreeNode) {
	// TODO
	node.Value = node.Left.Value + node.Right.Value
}

func (st *SegmentTree) query(node *SegmentTreeNode, start int, end int, left int, right int) int {
	// 区间 [l ,r] 完全包含区间 [start, end]
	// 例如：[2, 4] = [2, 2] + [3, 4]，当 [start, end] = [2, 2] 或者 [start, end] = [3, 4]，直接返回
	if left <= start && end <= right {
		return node.Value
	}
	// 把当前区间 [start, end] 均分得到左右孩子的区间范围
	// node 左孩子区间 [start, mid]
	// node 左孩子区间 [mid + 1, end]
	mid := (start + end) >> 1
	ans := 0
	// 下推标记
	st.pushDown(node, mid-start+1, end-mid)
	// [start, mid] 和 [l, r] 可能有交集，遍历左孩子区间
	if left <= mid {
		// TODO
		ans += st.query(node.Left, start, mid, left, right)
	}
	// [mid + 1, end] 和 [l, r] 可能有交集，遍历右孩子区间
	if right > mid {
		// TODO
		ans += st.query(node.Right, mid+1, end, left, right)
	}
	// ans 把左右子树的结果都累加起来了，与树的后续遍历同理
	return ans
}

func (st *SegmentTree) Update(left int, right int, value int) {
	st.update(st.Head, 0, st.Size-1, left, right, value)
}

func (st *SegmentTree) Query(left int, right int) int {
	return st.query(st.Head, 0, st.Size-1, left, right)
}
