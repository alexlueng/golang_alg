package calender

import "math"

// feature 5 given the schedules of two users, find the time intervals when both of them are busy

func meetingsIntersection(meetingA, meetingB [][]int) [][]int {
	i := 0
	j := 0

	var intersection [][]int
	for i < len(meetingA) && j < len(meetingB) {
		start := int(math.Max(float64(meetingA[i][0]), float64(meetingB[j][0])))
		end := int(math.Min(float64(meetingA[i][1]), float64(meetingB[j][1])))
		if start < end {
			intersection = append(intersection, []int{start, end})
		}
		if meetingA[i][1] < meetingB[j][1] {
			i++
		} else {
			j++
		}
	}
	return intersection
}
