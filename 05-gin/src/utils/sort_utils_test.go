package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCaseScenario(t *testing.T) {
	// initialisation section
	// creates a slice where every element will require processing/swapping
	eles := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	// execution section
	eles = BubbleSort(eles)

	// validation section
	// tests return is not nil
	assert.NotNil(t, eles)

	// check length of return is equal to input
	assert.EqualValues(t, 9, len(eles))

	// checks individual elements against
	// expected sort ascending
	assert.EqualValues(t, 1, eles[0])
	assert.EqualValues(t, 2, eles[1])
	assert.EqualValues(t, 3, eles[2])
	assert.EqualValues(t, 4, eles[3])
	assert.EqualValues(t, 5, eles[4])
	assert.EqualValues(t, 6, eles[5])
	assert.EqualValues(t, 7, eles[6])
	assert.EqualValues(t, 8, eles[7])
	assert.EqualValues(t, 9, eles[8])

}

func TestBubbleSortBestCaseScenario(t *testing.T) {
	eles := []int{1, 2, 3, 4, 5}
	eles = BubbleSort(eles)

	assert.NotNil(t, eles)
	assert.EqualValues(t, 5, len(eles))

	// checks individual elements against
	assert.EqualValues(t, 1, eles[0])
	assert.EqualValues(t, 2, eles[1])
	assert.EqualValues(t, 3, eles[2])
	assert.EqualValues(t, 4, eles[3])
	assert.EqualValues(t, 5, eles[4])
}

func TestBubbleSortNilCaseScenario(t *testing.T) {
	eles := BubbleSort(nil)
	assert.Nil(t, eles)
}

func getElements(n int) []int {
	res := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		res[i] = j
		i++
	}
	return res
}

func BenchmarkBubbleSort1000(b *testing.B) {
	eles := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(eles)

	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	eles := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(eles)

	}
}

func BenchmarkSort1000(b *testing.B) {
	eles := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(eles)

	}
}

func BenchmarkSort100000(b *testing.B) {
	eles := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(eles)

	}
}
