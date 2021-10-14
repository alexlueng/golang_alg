package calender

// feature 7 given the schedule of a person, find the longest consecutive busy period

func lengthOfLongestBusyPeriod(schedule []int) int {
	longestBusyPeriod := 0
	slotSet := make(map[int]bool)
	for _, i := range schedule {
		slotSet[i] = true
	}

	for slot, _ := range slotSet {
		if _, ok := slotSet[slot-1]; !ok {
			currentSlot := slot
			currentConsecutiveSlot := 1
			_, temp := slotSet[currentSlot+1]
			for temp {
				currentSlot++
				currentConsecutiveSlot++
				_, temp = slotSet[currentSlot+1]
			}
			if currentConsecutiveSlot > longestBusyPeriod {
				longestBusyPeriod = currentConsecutiveSlot
			}
		}
	}
	return longestBusyPeriod
}
