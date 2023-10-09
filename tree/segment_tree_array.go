package tree

type SegmentTreeArray []int

func NewSegmentTreeArray(nums []int) *SegmentTreeArray {
	str := make(SegmentTreeArray, len(nums)*4)
	str.build(0, 0, len(nums)-1, nums)
	return &str
}

func (str SegmentTreeArray) Len() int {
	return len(str) / 4
}

func (str SegmentTreeArray) build(node, s, e int, nums []int) {
	if s == e {
		str[node] = nums[s]
		return
	}
	m := s + ((e - s) >> 1)
	str.build(node*2+1, s, m, nums)
	str.build(node*2+2, m+1, e, nums)
	str[node] = str[node*2+1] + str[node*2+2]
}

func (str SegmentTreeArray) update(node, s, e, index, value int) {
	if s == e {
		str[node] = value
		return
	}
	m := s + ((e - s) >> 1)
	if index <= m {
		str.update(node*2+1, s, m, index, value)
	} else {
		str.update(node*2+2, m+1, e, index, value)
	}
	str[node] = str[node*2+1] + str[node*2+2]
}

func (str SegmentTreeArray) Update(index, value int) {
	str.update(0, 0, str.Len()-1, index, value)
}

func (str SegmentTreeArray) sumrange(node, s, e, left, right int) int {
	if s == left && e == right {
		return str[node]
	}
	m := s + ((e - s) >> 1)
	if right <= m {
		return str.sumrange(node*2+1, s, m, left, right)
	}
	if left > m {
		return str.sumrange(node*2+2, m+1, e, left, right)
	}
	return str.sumrange(node*2+1, s, m, left, m) + str.sumrange(node*2+2, m+1, e, m+1, right)
}

func (str SegmentTreeArray) SumRange(left, right int) int {
	return str.sumrange(0, 0, str.Len()-1, left, right)
}
