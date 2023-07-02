package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	left := items[:middle]
	right := items[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func main() {
	items := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:")
	fmt.Println(items)

	sortedItems := mergeSort(items)
	fmt.Println("Sorted array:")
	fmt.Println(sortedItems)
}
