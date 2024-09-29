package trie

// all will be stored as start_idx : 0, end_idx: 2. so length - e-s+1
type Location struct {
	KeywordIndex int
	StartIdx     int
	EndIdx       int
}

type Node struct {
	Location *Location
	Children []*Node
}

type Trie struct {
	Words []string
	Root  *Node
}

func NewTrie() *Trie {
	return &Trie{
		Words: []string{},
		Root:  &Node{},
	}
}

func (n *Node) GetLocation() *Location {
	return n.Location
}

func (n *Node) GetWord(word string, idx int, allWords *[]string) bool {
	if idx == len(word) {
		return true
	}
	// go through the children, see if the words match, if they do, go deeper.
	for _, child := range n.Children {
		isMatch := true
		childLength := child.Location.EndIdx - child.Location.StartIdx + 1
		// check if the remaining word is still withing range?
		if idx+childLength > len(word) {
			isMatch = false
			continue
		}
		// check if the word matches
		for i := 0; i < childLength; i++ {
			if word[idx+i] != (*allWords)[child.Location.KeywordIndex][child.Location.StartIdx+i] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return child.GetWord(word, idx+childLength, allWords)
		}
	}
	return false
}

func (n *Node) splitNode(cursor int, words *[]string, word string) {
	originalStartIdx := n.Location.StartIdx
	originalEndIdx := n.Location.EndIdx
	n.Location = &Location{
		KeywordIndex: n.Location.KeywordIndex,
		StartIdx:     originalStartIdx,
		EndIdx:       originalStartIdx + cursor - 1,
	}

	newNode := &Node{
		Location: &Location{
			KeywordIndex: len(*words) - 1,
			StartIdx:     cursor,
			EndIdx:       len(word) - 1,
		},
		Children: make([]*Node, 0),
	}
	siblingNodeChildren := make([]*Node, 0)
	for _, c := range n.Children {
		siblingNodeChildren = append(siblingNodeChildren, c)
	}
	siblingNode := &Node{
		Location: &Location{
			KeywordIndex: n.Location.KeywordIndex,
			StartIdx:     originalStartIdx + cursor,
			EndIdx:       originalEndIdx,
		},
		Children: siblingNodeChildren,
	}
	n.Children = []*Node{newNode, siblingNode}
}

// If any of the children matches, then iterate until it matches and then go in.
// If none matches, create a new node and add it to the children.
func (n *Node) AddWord(word string, idx int, words *[]string) {
	// go through the children, see if the words match, if they do, go deeper.
	for _, child := range n.Children {
		isMatch := true
		childLength := child.Location.EndIdx - child.Location.StartIdx + 1

		// check if the word matches completely
		cursor := 0
		for i := 0; i < childLength; i++ {
			if word[idx+i] != (*words)[child.Location.KeywordIndex][child.Location.StartIdx+i] {
				isMatch = false
				if cursor > 0 {
					// partial match, take the data till match, create a new node and add it to the children.
					// the rest of the data becomes a separate children.
					child.splitNode(cursor, words, word)
					return
				}
				break
			}
			cursor++
		}
		if isMatch {
			child.AddWord(word, idx+childLength, words)
			return
		}
	}
	// if none of the children match, then create a new node and add it to the children.
	newNode := &Node{
		Location: &Location{
			KeywordIndex: len(*words) - 1,
			StartIdx:     idx,
			EndIdx:       len(word) - 1,
		},
		Children: make([]*Node, 0),
	}
	n.Children = append(n.Children, newNode)
}

func (n *Node) MaxDepth() int {
	if len(n.Children) == 0 {
		return 0
	}
	maxChildrenDepth := 0
	for _, child := range n.Children {
		depth := child.MaxDepth()
		if depth > maxChildrenDepth {
			maxChildrenDepth = depth
		}
	}
	return maxChildrenDepth + 1
}

func (t *Trie) AddWord(word string) {
	if t.Root == nil {
		t.Root = &Node{}
	}
	if t.FindWord(word, 0) {
		return
	}
	t.Words = append(t.Words, word)
	t.Root.AddWord(word, 0, &t.Words)
}

func (t *Trie) FindWord(word string, idx int) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.GetWord(word, idx, &t.Words)
}
