package plagiarism

// Match a student’s code with the rest of the class’s code samples to identify the number of possible matches.
// Solution

func possibleMatches(s string, words []string) int {
	ans := 0
	waitingList := make(map[byte][]string)

	for _, w := range words {
		waitingList[byte(w[0])] = append(waitingList[byte(w[0])], w)
	}

	for _, c := range s {
		advance := waitingList[byte(c)]
		waitingList[byte(c)] = []string{}
		for _, it := range advance {
			if len(it) >= 2 {
				waitingList[byte(it[1])] = append(waitingList[byte(it[1])], it[1:])
			} else {
				ans++
			}
		}
	}

	return ans
}
