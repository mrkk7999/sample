package main

import "fmt"

func BubbleSortNaive(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		var flag = 0
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
}
func BubbleSortOptimal(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		var flag = 0
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
}
func BubbleSortRecursive(arr *[]int, n int) {
	if n == 1 {
		return
	}
	var count = 0
	for i := 0; i < n-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
			count = count + 1
		}
	}
	if count == 0 {
		return
	}
	BubbleSortRecursive(arr, n-1)
}

func main() {
	arr := []int{32, 34, 65, 12, 11, 2, 1}
	//BubbleSortNaive(&arr)
	//BubbleSortNaive(&arr)
	BubbleSortRecursive(&arr, len(arr))
	fmt.Println(arr)
}
