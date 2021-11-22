package zoom

// feature 1 implement pagination for the display showing the names of participants present in a meeting

type DisplayLobby struct {
	stack Stack
}

func new(root *Node) *DisplayLobby {
	lobby := &DisplayLobby{}
	lobby.PushAll(root)
	return lobby
}

func (l *DisplayLobby) PushAll(node *Node) {
	for node != nil {
		l.stack = l.stack.Push(node)
		node = node.leftChild
	}
}

func (l *DisplayLobby) HasNext() bool {
	return !l.stack.Empty()
}

func (l *DisplayLobby) NextName() string {
	var tmpNode *Node
	l.stack, tmpNode = l.stack.Pop()
	l.PushAll(tmpNode.rightChild)
	return tmpNode.val
}

func (l *DisplayLobby) NextPage() []string {
	var res []string
	for i := 0; i < 10; i++ {
		if l.HasNext() {
			res = append(res, l.NextName())
		} else {
			break
		}
	}
	return res
}
