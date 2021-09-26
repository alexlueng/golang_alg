package facebook

// feature 3 sync the facebook stories list with instagram

func search(arr []int, start, end, key int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if arr[mid] == key {
		return mid
	}
	if arr[start] <= arr[mid] && key <= arr[mid] && key >= arr[start] {
		return search(arr, start, mid-1, key)
	} else if arr[mid] <= arr[end] && key >= arr[mid] && key <= arr[end] {
		return search(arr, mid+1, end, key)
	} else if arr[end] <= arr[mid] {
		return search(arr, mid+1, end, key)
	} else if arr[start] >= arr[mid] {
		return search(arr, start, mid-1, key)
	}
	return -1
}

func findStoryID(arr []int, key int) int {
	return search(arr, 0, len(arr)-1, key)
}
