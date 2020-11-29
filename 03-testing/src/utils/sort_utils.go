package utils

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
