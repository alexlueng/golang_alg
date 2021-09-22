package netflix

// feature 7 During a user session, a user often "shops" around for a program to watch. During this session,
// we want to let them move back and forth in the history of programs they've just browsed.
// As a developer, you can smell a stack, right? But, we also want the user to be able to derictly jump
// to the top-ranked program from the one's they've browsed.

// Solution
// Use mainstack and maxstack
// maxstack checks the value being pushed. If maxstack is empty, this value is pushed into it and becomes the current maximum. If maxstack already has elements in it, the value is compared with the top element in this stack. The element is inserted into maxstack if it is greater than the top element

type MaxStack struct {
	maxSize   int
	mainStack Stack
	maxStack  Stack
}

func (m *MaxStack) Pop() int {
	m.maxStack, _ = m.maxStack.Pop()
	top := m.mainStack.Top()
	m.mainStack, _ = m.mainStack.Pop()
	return top
}

func (m *MaxStack) Push(value int) {
	m.mainStack = m.mainStack.Push(value)
	if !m.maxStack.Empty() && m.maxStack.Top() > value {
		m.maxStack = m.maxStack.Push(m.maxStack.Top())
	} else {
		m.maxStack.Push(value)
	}
}

func (m *MaxStack) MaxRating() int {
	return m.maxStack.Top()
}

// feature 8 a user complained that the next and previous functionality isn't working correctly.
// Using their session history, we want to check if our implementation is correct or indeed buggy

func verifySession(pushOp, popOp []int) bool {
	M := len(pushOp)
	N := len(popOp)
	if M != N {
		return false
	}
	var stack Stack
	i := 0
	for _, pid := range pushOp {
		stack = stack.Push(pid)
		for !stack.Empty() && stack.Top() == popOp[i] {
			stack, _ = stack.Pop()
			i++
		}
	}
	return stack.Empty()
}
