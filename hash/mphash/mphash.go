// multi-probe consistent hash
package mphash

import (
	"crypto/md5"
	"encoding/binary"
	"sort"
)

type Node struct {
	Id     string
	Weight int
}

type Ring struct {
	Nodes       []Node
	NumReplicas int
	Probes      int
}

func NewRing(numReplicas, probes int, nodes []Node) *Ring {
	return &Ring{
		Nodes:       nodes,
		NumReplicas: numReplicas,
		Probes:      probes,
	}
}

func (r *Ring) AddNode(node Node) {
	r.Nodes = append(r.Nodes, node)
}

func (r *Ring) RemoveNode(node Node) {
	for i, n := range r.Nodes {
		if n.Id == node.Id {
			r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)
			break
		}
	}
}

func (r *Ring) GetNode(key string) Node {
	if len(r.Nodes) == 0 {
		return Node{}
	}

	// Create a set of probing points
	probingPoints := r.getProbingPoints(key)

	// Sort the nodes by their hash value
	sort.Slice(r.Nodes, func(i, j int) bool {
		return r.hash(r.Nodes[i].Id) < r.hash(r.Nodes[j].Id)
	})

	// Initialize the candidate set of nodes with the first node
	candidates := []Node{r.Nodes[0]}

	// Search for the closest set of nodes that collectively handle the key
	for i := 1; i < len(r.Nodes) && len(candidates) < r.Probes; i++ {
		node := r.Nodes[i]

		// Check if the node is in the range of any of the probing points
		for _, point := range probingPoints {
			if r.hash(node.Id)-point >= r.hash(candidates[len(candidates)-1].Id)-point {
				break
			}
			if !r.containsNode(candidates, node) {
				candidates = append(candidates, node)
				break
			}
		}
	}

	// Return the first node in the candidate set
	return candidates[0]
}

func (r *Ring) containsNode(nodes []Node, node Node) bool {
	for _, n := range nodes {
		if n.Id == node.Id {
			return true
		}
	}
	return false
}

func (r *Ring) hash(key string) uint32 {
	sum := md5.Sum([]byte(key))
	return binary.LittleEndian.Uint32(sum[:4])
}

func (r *Ring) getProbingPoints(key string) []uint32 {
	probingPoints := make([]uint32, r.Probes)
	hash := r.hash(key)
	for i := 0; i < r.Probes; i++ {
		probingPoints[i] = (hash + uint32(i)*hash) % ((1 << 32) - 1)
	}
	return probingPoints
}
