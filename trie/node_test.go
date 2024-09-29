package trie

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// create a test to get a word from the trie
func TestShouldGetWordSuccessfully(t *testing.T) {
	// l1 : he, world
	// l2: llo, ck
	l2Nodes := []*Node{
		{
			Location: &Location{
				KeywordIndex: 0,
				StartIdx:     2,
				EndIdx:       4,
			},
			Children: make([]*Node, 0),
		},
		{
			Location: &Location{
				KeywordIndex: 2,
				StartIdx:     2,
				EndIdx:       3,
			},
			Children: make([]*Node, 0),
		},
	}
	l1Nodes := []*Node{
		{
			Location: &Location{
				KeywordIndex: 0,
				StartIdx:     0,
				EndIdx:       1,
			},
			Children: l2Nodes,
		},
		{
			Location: &Location{
				KeywordIndex: 1,
				StartIdx:     0,
				EndIdx:       4,
			},
			Children: make([]*Node, 0),
		},
	}
	trie := &Trie{
		Words: []string{"hello", "world", "heck"},
		Root: &Node{
			Location: nil,
			Children: l1Nodes,
		},
	}

	assert.True(t, trie.Root.GetWord("hello", 0, &trie.Words))
	assert.False(t, trie.Root.GetWord("helo", 0, &trie.Words))
	assert.True(t, trie.Root.GetWord("world", 0, &trie.Words))
	assert.True(t, trie.Root.GetWord("heck", 0, &trie.Words))
}

func TestShouldAddWordInEmptyTrie(t *testing.T) {
	trie := &Trie{}
	trie.AddWord("hello")
	assert.Equal(t, 1, len(trie.Words))
	assert.Equal(t, "hello", trie.Words[0])
	assert.True(t, trie.FindWord("hello", 0))
}

func TestShouldAddMultipleWordsInTrie(t *testing.T) {
	trie := &Trie{}
	testWords := []string{"hello", "world", "heck", "hakuna", "wordpad"}
	expectedDepths := []int{1, 1, 2, 3, 3}
	for i, word := range testWords {
		trie.AddWord(word)
		assert.Equal(t, expectedDepths[i], trie.Root.MaxDepth())
	}
	assert.Equal(t, 5, len(trie.Words))
	for _, word := range testWords {
		assert.True(t, trie.FindWord(word, 0))
	}
}
