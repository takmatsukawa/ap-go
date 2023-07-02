package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap arr[i] and arr[minIdx]
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:")
	fmt.Println(arr)

	selectionSort(arr)
	fmt.Println("Sorted array:")
	fmt.Println(arr)
}
