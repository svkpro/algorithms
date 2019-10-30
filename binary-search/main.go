package main

import "fmt"

func main() {
	items := []int{0, 1, 2, 3, 8, 11, 21}
	fmt.Println(binarySearch(8, items))
}

func binarySearch(needle int, haystack []int) bool {

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
		return false
	}

	return true
}
