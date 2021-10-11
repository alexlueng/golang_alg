package searchengine

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Empty() bool        { return len(h) == 0 }
func (h MaxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() []int         { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
