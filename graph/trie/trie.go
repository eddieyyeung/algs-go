package trie

type Node struct {
	Children    map[int]*Node
	IsEndOfWord bool
}

func New() *Node {
	return &Node{
		Children: map[int]*Node{},
	}
}

func (root *Node) Insert(key string) {
	p := root
	for _, char := range key {
		if _, ok := p.Children[int(char)]; !ok {
			p.Children[int(char)] = &Node{
				Children:    map[int]*Node{},
				IsEndOfWord: false,
			}
		}
		p = p.Children[int(char)]
	}
	p.IsEndOfWord = true
}

func (root *Node) Search(key string) bool {
	p := root
	for _, char := range key {
		if _, ok := p.Children[int(char)]; !ok {
			return false
		}
		p = p.Children[int(char)]
	}
	return p.IsEndOfWord
}
