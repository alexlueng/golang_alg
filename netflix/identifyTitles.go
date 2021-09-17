package netflix

// Feature 4 for efficiently distributing content to different geographic regions and for program recommendation to viewers, we want to determine titles that are gaining or losing popularity scores
// We maintains a popularity score for each of its titles which derived from customer feedback, likes, dislikes, etc. We'll be provided with a list of integers representing the popularity scores of a movie collected over a number of weeks.

// Solution
// a list is increasing if the expression list[i] <= list[i+1]

func identifyTitles(scores []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(scores)-1; i++ {
		if scores[i] > scores[i+1] {
			increasing = false
		}
		if scores[i] < scores[i+1] {
			decreasing = false
		}
	}
	return increasing || decreasing
}
