package plagiarism

// feature 2 match code samples of two students to identify which part has been copied and altered
// to avoid plagiarism detection.
// Solution
// 1) traverse the cheater string, and for each letter, check if it's equal to the current letter in student
// 2) if the letters are equal, move to the next letter in both the strings.
// 3) when the last letter of student is matched with a letter of cheater, mark that letter position of
// cheater as a possible window
// 4) go backward from the end of student and match the letters with cheater from the marked position.
// do this until the student string is exhausted
// 5) check if your current window is smaller than the previous one. if yes, mark it as the minimum window
// 6) repeat from step 2 until you reach the end of cheater

func solution(cheater, student string) string {
	window := ""
	j := 0
	min := len(cheater) + 1
	for i := 0; i < len(cheater); i++ {
		if cheater[i] == student[j] {
			j++
			if j == len(student) {
				end := i + 1
				j--
				for j >= 0 {
					if cheater[i] == student[j] {
						j--
					}
					i--
				}
				j++
				i++
				if end-i < min {
					min = end
					window = cheater[i:min]
				}
			}

		}
	}
	return window
}
