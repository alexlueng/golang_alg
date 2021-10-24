package uber

import "container/heap"

// feature 1 select at least the n nearest drivers within the user's vicinity, avoiding the drivers
// that are tow far away.
// Solution
// The dest data structure that comes to mind to track the nearest K drivers is Heap.

type Location struct {
	x, y int
}

func (l Location) distFromOrigin() int {
	return (l.x * l.x) + (l.y * l.y)
}

func newHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

func findClosestDrivers(locations []Location, k int) []Location {
	maxHeap := newHeap()
	var res []Location
	for i := 0; i < k; i++ {
		heap.Push(maxHeap, locations[i])
	}

	for i := k; i < len(locations); i++ {
		top := maxHeap.Top()
		if locations[i].distFromOrigin() < top.distFromOrigin() {
			maxHeap.Pop()
			heap.Push(maxHeap, locations[i])
		}
	}

	for !maxHeap.Empty() {
		res = append(res, maxHeap.Top())
		heap.Pop(maxHeap)
	}
	return res
}
