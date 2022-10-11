package main

func SelectionSort(arr []int) []int {
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		minIdx := i
		for j := i; j < arrLength; j++ {
			if arr[j] > arr[minIdx] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
	return arr
}
func main() {
	//var arr = []int{21, 34, 76, 87, 23, 1, 2}
	//arr = SelectionSort(arr)
	//fmt.Println(arr)
	var res [][]int
	temp := []int{1, 2, 3}
	res = append(res, temp, temp)
}
