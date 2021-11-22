package network

// feature 6 A pair of machines have a request-response message exchange. Identify the request and response
// packets used the same path.

func isPalindrome(arr []int, left, right int) bool {
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func transmissionError(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left] == arr[right] {
			left++
			right--
		} else {
			if isPalindrome(arr, left+1, right) {
				return 1
			}
			if isPalindrome(arr, left, right-1) {
				return 1
			}
			return -1
		}
	}
	return 0
}
