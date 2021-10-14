package calender

import (
	"container/heap"
	"sort"
)

// feature 1 We have a schedule of meetings available. We want to determine and block the minimum
// number of meeting rooms for these meetings.
// Our job is to implement a solution that can identify the umber of meeting rooms needed to
// schedule the required meetings.
// Solution
// We can use a heap or priority queue to keep the meeting timings. If the room we obtain from the top of
// the heap isn't free yet, then this means no other room is free either. So we allocate a new room in this
// condition.
// 1) sort the given meetings by their startTime
// 2) allocate the first meeting to a room and add an entry in the heap with the meeting's endTime
// 3) traverse over the remaining meetings, and keep checking if the meeting at the top of the heap has
// ended or not. This will tell us if a room is free yet.
// 4) if the room is free, then we extract this element and it again in the heap with the ending time of
// the current meeting we are processing.
// 5) if not, then we allocate a new room and add it to the heap
// 6) after processing all the meetings, the heap's size will tell us the number of rooms allocated. This
// will be the minimum number of rooms needed to accommodate all the meetings.

func minMeetingRooms(meetingTimes [][]int) int {
	if len(meetingTimes) == 0 {
		return 0
	}
	// 1)
	sort.Slice(meetingTimes, func(i, j int) bool {
		return meetingTimes[i][0] < meetingTimes[j][0]
	})

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, meetingTimes[0][1])

	for i := 1; i < len(meetingTimes); i++ {
		currStart := meetingTimes[i][0]
		currEnding := meetingTimes[i][1]
		earliestEnding := minHeap.Top()

		if earliestEnding <= currStart {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, currEnding)
	}
	return minHeap.Len()
}
