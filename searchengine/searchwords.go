package searchengine

// Feature 1 Design a system that can store and fetch words efficiently. This can be used to store
// web pages to make searching easier
// Solution
// We will use a trie data structure.
// Not using hash table: it would be very inefficient for the prefix searching. And scaling hashtables
// for large datasets also increases collisions and space complexity.

type Node struct {
	children map[rune]*Node
	isWord   bool
}

type WordDictionary struct {
	root *Node
}

func new() WordDictionary {
	node := &Node{
		children: make(map[rune]*Node),
	}
	return WordDictionary{root: node}
}

func (w *WordDictionary) InsertWord(word string) {
	node := w.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &Node{children: make(map[rune]*Node)}
		}
		node = node.children[c]
	}
	node.isWord = true
}

func (w *WordDictionary) SearchWord(word string) bool {
	node := w.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return node.isWord
}
