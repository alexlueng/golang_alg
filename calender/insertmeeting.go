package calender

import (
	"math"
	"sort"
)

// feature 4 some users have open agenda and open participant meetings so that two meetings cant take
// place at overlapping time at the same venue. Add a new meeting to a user's already busy schedule
// while merging the new meeting with existing meetings, if necessary.

func insertMeeting(meetingTimes [][]int, newMeeting []int) [][]int {
	var output [][]int
	sort.Slice(meetingTimes, func(i, j int) bool {
		return meetingTimes[i][0] < meetingTimes[j][0]
	})
	i := 0
	n := len(meetingTimes)
	for i < n && newMeeting[0] > meetingTimes[i][0] {
		output = append(output, meetingTimes[i])
		i++
	}
	size := len(output)
	if size == 0 || output[size-1][1] < newMeeting[0] {
		output = append(output, newMeeting)
	} else {
		output[size-1][1] = int(math.Max(float64(output[size-1][1]), float64(newMeeting[1])))
	}

	for i < n {
		size = len(output)
		if output[size-1][1] < meetingTimes[i][0] {
			output = append(output, meetingTimes[i])
		} else {
			output[size-1][1] = int(math.Max(float64(output[size-1][1]), float64(meetingTimes[i][1])))
		}

	}
	return output
}
