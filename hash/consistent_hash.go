package hash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// HashRing struct represents a consistent hash ring
type HashRing struct {
	nodes       map[uint32]string
	sortedNodes []uint32
}

// NewHashRing creates a new HashRing
func NewHashRing() *HashRing {
	return &HashRing{
		nodes: make(map[uint32]string),
	}
}

// AddNode adds a node to the HashRing
func (hr *HashRing) AddNode(node string) {
	hash := crc32.ChecksumIEEE([]byte(node))
	hr.nodes[hash] = node
	hr.sortedNodes = append(hr.sortedNodes, hash)
	sort.Slice(hr.sortedNodes, func(i, j int) bool {
		return hr.sortedNodes[i] < hr.sortedNodes[j]
	})
}

// RemoveNode removes a node from the HashRing
func (hr *HashRing) RemoveNode(node string) {
	hash := crc32.ChecksumIEEE([]byte(node))
	delete(hr.nodes, hash)
	for i := range hr.sortedNodes {
		if hr.sortedNodes[i] == hash {
			hr.sortedNodes = append(hr.sortedNodes[:i], hr.sortedNodes[i+1:]...)
			break
		}
	}
}

// GetNode returns the node responsible for a given key
func (hr *HashRing) GetNode(key string) string {
	hash := crc32.ChecksumIEEE([]byte(key))
	i := sort.Search(len(hr.sortedNodes), func(i int) bool {
		return hr.sortedNodes[i] >= hash
	})
	if i == len(hr.sortedNodes) {
		i = 0
	}
	return hr.nodes[hr.sortedNodes[i]]
}

func main() {
	hr := NewHashRing()
	hr.AddNode("node1")
	hr.AddNode("node2")
	hr.AddNode("node3")

	for i := 0; i < 10; i++ {
		key := "key" + strconv.Itoa(i)
		node := hr.GetNode(key)
		println(key, "->", node)
	}

	hr.RemoveNode("node2")

	for i := 0; i < 10; i++ {
		key := "key" + strconv.Itoa(i)
		node := hr.GetNode(key)
		println(key, "->", node)
	}
}
