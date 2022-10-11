package main

import "fmt"

//	type ListNode struct {
//		Val  int
//		Next *ListNode
//	}
//
//	func binaryToDecimal(binary []int) int {
//		decimal := 0
//		index := 0
//		remainder := 0
//		for i := 0; i < len(binary); i++ {
//			remainder = binary[i] % 10
//			decimal = decimal + remainder*math.Pow(2, float(index))
//			index++
//		}
//		return decimal
//	}
//
//	func getDecimalValue(head *ListNode) int {
//		var binary []int
//		temp := head
//		for temp != nil {
//			binary = append(binary, temp.Val)
//			temp = temp.Next
//		}
//		if len(binary) == 1 {
//			return binary[0]
//		}
//		return binaryToDecimal(binary)
//	}
func LinearSearch(arr []int, target int) (int, bool) {
	if len(arr) > 0 && arr != nil {
		var i = 0
		for i < len(arr) {
			if target == arr[i] {
				return i, true
			}
		}
	}
	return -1, false
}

func BinarySearchIterative(arr []int, target int) (int, bool) {
	low, high := 0, len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		if target == arr[mid] {
			return mid, true
		}
		if target > arr[mid] {
			low = mid + 1
		} else if target < arr[mid] {
			high = mid - 1
		}
	}
	return -1, false
}

func BinarySearchRecursive(arr []int, target, low, high int) (int, bool) {
	// Base Case
	if low > high {
		return -1, false
	}
	// calculating mid
	mid := low + (high-low)/2
	// checking where the element lie

	// element in second half
	if target > arr[mid] {
		return BinarySearchRecursive(arr, target, mid+1, high)
	} else if target < arr[mid] {
		// first half
		return BinarySearchRecursive(arr, target, low, mid-1)
	} else {
		// at the mid position
		return mid, true
	}
	// Not found
}

func main() {
	arr := []int{21, 34, 65, 87, 23}
	target := 21
	//idx, isFound := LinearSearch(arr, target)
	//if isFound {
	//	fmt.Printf("Element : %d found at position : %d", target, idx)
	//} else {
	//	fmt.Println("Not present")
	//}
	//idx, isFound = BinarySearchIterative(arr, target)
	//if isFound {
	//	fmt.Printf("Element : %d found at position : %d", target, idx)
	//} else {
	//	fmt.Println("Not present")
	//}
	low, high := 0, len(arr)-1
	idx, ok := BinarySearchRecursive(arr, target, low, high)
	if ok {
		fmt.Printf("Element : %d found at position : %d", target, idx)
	} else {
		fmt.Println("Not present")
	}
}
