package sort

import "fmt"

var (
	swapCount       int
	swapRealCount   int
	partitionsCount int
	quickSortCount  int
)

func statOutput() {
	fmt.Printf("Statistics:\n")
	fmt.Printf("\tsort calls      %d\n", quickSortCount)
	fmt.Printf("\tpart calls      %d\n", partitionsCount)
	fmt.Printf("\tswap calls      %d\n", swapCount)
	fmt.Printf("\treal swap calls %d\n", swapRealCount)
}

func swap(a, b *byte) {
	swapCount++
	if a == b {
		return
	}
	swapRealCount++

	temp := *a
	*a = *b
	*b = temp
}

func partition(array []byte, low, high int) int {
	partitionsCount++
	p := high
	firsthigh := low
	for i := low; i < high; i++ {
		if array[i] < array[p] {
			swap(&array[i], &array[firsthigh])
			firsthigh++
		}
	}
	swap(&array[p], &array[firsthigh])
	return firsthigh
}

func quicksort(array []byte, low, high int) {
	quickSortCount++
	if high-low > 0 {
		p := partition(array, low, high)
		quicksort(array, low, p-1)
		quicksort(array, p+1, high)
	}
}

func SkienaSort(array []byte) {
	quicksort(array, 0, len(array)-1)
	statOutput()
}
