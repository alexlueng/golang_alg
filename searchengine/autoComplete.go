package searchengine

import (
	"sort"
)

// Feature 2 Implement the auto-complete functionality to apply when the user is entering a query.
// For this feature, you will be recommending auto-completion using a set of popular, already-available queries

// we will again use the trie data structure.

type ANode struct {
	children map[rune]*ANode
	isEnd    bool
	data     string
	rank     int
	hot      []*ANode
}

func (a *ANode) Update(n *ANode) {
	nd := sort.Search(len(a.hot), func(i int) bool { return a.hot[i] == n })
	if nd <= len(a.hot) {
		a.hot = append(a.hot, n)
	}
	sort.Slice(a.hot, func(i, j int) bool {
		return a.hot[i].rank > a.hot[j].rank
	})
	if len(a.hot) > 3 {
		a.hot = a.hot[:len(a.hot)-1]
	}
}

type AutoCompleteSystem struct {
	root    *ANode
	current *ANode
	keyword string
}

func newAutoCompleteSystem(sentences []string, times []int) *AutoCompleteSystem {
	r := &ANode{children: make(map[rune]*ANode)}
	auto := &AutoCompleteSystem{
		root:    r,
		current: r,
	}
	for i := 0; i < len(sentences); i++ {
		auto.AddRecord(sentences[i], times[i])
	}
	return auto
}

func (a *AutoCompleteSystem) AddRecord(sentence string, t int) {
	node := a.root
	var visited []*ANode

	for _, c := range sentence {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &ANode{children: make(map[rune]*ANode)}
		}
		node = node.children[c]
		visited = append(visited, node)
	}
	node.isEnd = true
	node.data = sentence
	node.rank += t
	for _, i := range visited {
		i.Update(node)
	}
}

func (a *AutoCompleteSystem) AutoComplete(c rune) []string {
	var res []string
	if c == '#' {
		a.AddRecord(a.keyword, 1)
		a.keyword = ""
		a.current = a.root
		return res
	}
	a.keyword += string(c)
	if a.current != nil {
		if _, ok := a.current.children[c]; !ok {
			return res
		} else {
			a.current = a.current.children[c]
		}
	} else {
		return res
	}

	for _, node := range a.current.hot {
		res = append(res, node.data)
	}
	return res
}
