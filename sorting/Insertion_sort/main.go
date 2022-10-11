package main

import "fmt"

func InsertionSortIterative(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		j := i
		for j > 0 {
			if (*arr)[j-1] > (*arr)[j] {
				(*arr)[j-1], (*arr)[j] = (*arr)[j], (*arr)[j-1]
			}
			j = j - 1
		}
	}
}
func InsertionSortRecursive(arr *[]int, n int) {
	if n <= 1 {
		return
	}
	InsertionSortRecursive(arr, n-1)
	var last = (*arr)[n-1]
	var j = n - 2
	for j >= 0 && (*arr)[j] > last {
		(*arr)[j+1] = (*arr)[j]
		j--
	}
	(*arr)[j+1] = last
}
func main() {
	arr := []int{76, 87, 3, 23, 65, 1}
	// InsertionSortIterative(&arr)
	fmt.Println(arr)
	InsertionSortRecursive(&arr, len(arr))
	fmt.Println(arr)
}
