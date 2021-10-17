package stockscraper

type Stack []*LCATreeNode

func (s Stack) Push(n *LCATreeNode) Stack {
	return append(s, n)
}

func (s Stack) Pop() (Stack, *LCATreeNode) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	return s[:l-1], s[l-1]
}

func (s Stack) Top() *LCATreeNode {
	l := len(s)
	return s[l-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}
