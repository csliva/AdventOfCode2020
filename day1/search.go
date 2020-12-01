package day1

//returns -1 for false, and an index position for true
func binarySearch(needle int, haystack []int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return -1
	}
	return low
}
