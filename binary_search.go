package search

// BinarySearch searches a in s. s must be ordered.
// If search succeeds, returns an index of s, otherwise returns -1.
func BinarySearch(s []int, a int) int {
	high := len(s)-1
	low := 0

	for low <= high {
		mid := (low+high)/2
		if s[mid] == a {
			return mid
		} else if s[mid] > a {
			high = mid - 1
		} else { // s[mid] < a
			low = mid + 1
		}
	}

	return -1
}
