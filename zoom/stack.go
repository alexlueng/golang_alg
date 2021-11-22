package zoom

type Stack []*Node

func (s Stack) Push(n *Node) Stack {
	return append(s, n)
}

func (s Stack) Pop() (Stack, *Node) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	return s[:l-1], s[l-1]
}

func (s Stack) Top() *Node {
	l := len(s)
	return s[l-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}
