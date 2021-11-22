package network

import "math"

// feature 10 determine the longest stretch of days for which a customer's traffic variation was within
// a specified threshold.

func NewDeque() *Deque {
	return &Deque{}
}

func miniVariationLength(nums []int, threshold int) int {
	maxDeque := NewDeque()
	minDeque := NewDeque()
	start, end, ans := 0, 0, 0

	for end < len(nums) {
		for !minDeque.Empty() && nums[end] < nums[minDeque.Back()] {
			minDeque.PopBack()
		}

		for !maxDeque.Empty() && nums[end] > nums[maxDeque.Front()] {
			maxDeque.PopBack()
		}

		minDeque.PushBack(end)
		maxDeque.PushBack(end)

		variation := nums[maxDeque.Front()] - nums[minDeque.Front()]
		if variation > threshold {
			start++
			if start > minDeque.Front() {
				minDeque.PopFront()
			}
			if start > maxDeque.Front() {
				maxDeque.PopFront()
			}
		}
		ans = int(math.Max(float64(ans), float64(end-start+1)))
		end++
	}
	return ans
}
