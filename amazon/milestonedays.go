package amazon

// feature 5 we want to showcase order-processing milestones on a dashboard. We need to mine daily stats
// to populate this dashboard
// Solution
// use a modified binary search approach.

func search(milestones []int, n int) int {
	first := 0
	last := len(milestones)
	for first < last {
		var mid int = (first + last) / 2
		if milestones[mid] >= n {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func milestoneDays(milestones []int, target int) (int, int) {
	firstDay := search(milestones, target)
	if target == milestones[firstDay] {
		lastDay := search(milestones, target+1) - 1
		return firstDay, lastDay
	}
	return -1, -1
}
