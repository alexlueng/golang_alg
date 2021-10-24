package uber

import "container/heap"

// feature 7 instead of always picking the top-ranked drivers, smartly pair drivers to customers so
// that high-ranked drivers are more likely, but others are also picked sometimes.

func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

func kthHighestRank(ranks []int, k int) int {
	minHeap := newMinHeap()

	for i := 0; i < k; i++ {
		heap.Push(minHeap, ranks[i])
	}

	for i := k; i < len(ranks); i++ {
		if ranks[i] > minHeap.Top() {
			heap.Pop(minHeap)
			heap.Push(minHeap, ranks[i])
		}
	}

	return minHeap.Top()
}
