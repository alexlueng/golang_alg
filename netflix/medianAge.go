package netflix

import "container/heap"

// as part of a demographic study, we are interested in the median age of our viewers. We want to implement a functionality whereby the median age can be updated efficiently whenever a new user signs up.

// Solution
// assume that 'x' is the median age of a user. Half of the ages in the list will be smaller than 'x'
// and half of will be greater. So we can divide the list into two halves: one half to store the smaller
// numbers and one half to store the larger numbers.
// the median of all ages will either be the largest number in the smallList or the smallest number in the
// largestList. If the total number of elements is even, we know that the median will be the average of
// these two numbers. The best data structure for this is a Heap
// 1 store the first half of the numbers in a max heap. and the other in a min heap
// 2 calculate the median age

type MedianOfAges struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func newMedianAge() *MedianOfAges {
	min := &MinHeap{}
	max := &MaxHeap{}
	heap.Init(min)
	heap.Init(max)
	return &MedianOfAges{minHeap: min, maxHeap: max}
}

func (m *MedianOfAges) FindMedian() float64 {
	if m.maxHeap.Len() == m.minHeap.Len() {
		return float64(float64(m.maxHeap.Top()+m.minHeap.Top()) / 2.0)
	}
	// max heap will have one more element than the min heap
	return float64(m.maxHeap.Top())
}

func (m *MedianOfAges) InsertNum(num int) {
	if m.maxHeap.Empty() || m.maxHeap.Top() >= num {
		heap.Push(m.maxHeap, num)
	} else {
		heap.Push(m.minHeap, num)
	}

	// either both the heaps will have equal number of elements or max-heap will have one
	// more element than the min-heap
	if m.maxHeap.Len() > m.minHeap.Len()+1 {
		heap.Push(m.minHeap, heap.Pop(m.maxHeap).(int))
	} else if m.maxHeap.Len() < m.minHeap.Len() {
		heap.Push(m.maxHeap, heap.Pop(m.minHeap).(int))
	}
}
