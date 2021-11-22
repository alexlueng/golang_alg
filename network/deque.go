package network

type Deque struct {
	Items []int
}

func (s *Deque) PushFront(item int) {
	temp := []int{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) PushBack(item int) {
	s.Items = append(s.Items, item)
}

func (s *Deque) PopFront() int {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) PopBack() int {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) Front() int {
	return s.Items[0]
}

func (s *Deque) Back() int {
	return s.Items[len(s.Items)-1]
}

func (s *Deque) Empty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}
