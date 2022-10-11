package main

import "fmt"

func Merge(leftArr, rightArr []int) []int {
	var i, j = 0, 0
	// Don't make this mistake again
	//res := make([]int, len(leftArr)+len(rightArr))
	var res []int
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			res = append(res, leftArr[i])
			i++
		} else {
			res = append(res, rightArr[j])
			j++
		}
	}
	for i < len(leftArr) {
		res = append(res, leftArr[i])
		i++
	}
	for j < len(rightArr) {
		res = append(res, rightArr[j])
		j++
	}
	return res
}

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := Sort(arr[:len(arr)/2])
	second := Sort(arr[len(arr)/2:])
	return Merge(first, second)
}

//func Merge(leftArr, rightArr []int) []int {
//
//}
//
//func Sort(arr []int) []int {
//	var low,high = 0,len(arr)
//	for low<high{
//		var mid = low + (high-low)/2
//		Sort(arr)
//	}
//}

func main() {
	arr := []int{76, 87, 3, 23, 65, 1}
	arr = Sort(arr)
	fmt.Println(arr)
}
