package calender

import (
	"math"
	"sort"
)

// feature 2 Some team leads have requested the ability to display the blocks of time during which
// they will be busy in meetings. This will help others who want to meet them efficiently request an
// appropriate meeting time as well as to avoid unnecessary walk-ins

func mergeMeetings(meetingTimes [][]int) [][]int {
	var merged [][]int
	sort.Slice(meetingTimes, func(i, j int) bool {
		return meetingTimes[i][0] < meetingTimes[j][0]
	})

	for _, meeting := range meetingTimes {
		size := len(merged)
		if size == 0 || merged[size-1][1] < meeting[0] {
			merged = append(merged, meeting)
		} else {
			merged[size-1][1] = int(math.Max(float64(merged[size-1][1]), float64(meeting[1])))
		}
	}
	return merged
}
