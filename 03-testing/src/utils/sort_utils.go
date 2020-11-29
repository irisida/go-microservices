package utils

import "sort"

// BubbleSort sorting algo
func BubbleSort(elements []int) []int {
	running := true

	for running {
		running = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				running = true
			}
		}
	}
	return elements
}

// SortEls takes in a slice of int and returns the sorted
// list by choosing the appropriate sort method based on
// the size of the elements list
func SortEls(elements []int) []int {

	if len(elements) < 2000 {
		return BubbleSort(elements)
	}
	sort.Ints(elements)
	return elements
}
