package facebook

// Feature 5 Identify the morphed versions of abused and profane words so posts containing them can
// be flagged inappropriate

// Solution
// 1) Initialize two pointers, i and j, to start traversing from S and W, respectively.
// 2) Check if letters currently pointed to by i and j of both words ar equal. Otherwise, return false
// 3) For each equal letter found:
// 		get the length of the repeating sequences of the equal letter found in both words
//		the length of the repeating sequence of W letters should be less than or equal to the length of
//		the repeating sequence of S letters
//		the length of the repeating sequence of S letters should be greater than or equal to 3
// 4) If any of the conditions mentioned in step 3 fails, return false

func repeatedLetters(s string, ind int) int {
	temp := ind
	for temp < len(s) && s[temp] == s[ind] {
		temp++
	}
	return temp - ind
}

func flagWords(S, W string) bool {
	if len(S) == 0 || len(W) == 0 {
		return false
	}
	i, j := 0, 0
	for i < len(S) && j < len(W) {
		if S[i] == W[j] {
			len1 := repeatedLetters(S, i)
			len2 := repeatedLetters(W, j)
			if len1 < 3 && len1 != len2 || len1 >= 3 && len1 < len2 {
				return false
			}
			i += len1
			j += len2
		} else {
			return false
		}
	}
	return i == len(S) && j == len(W)
}
