package main

import "fmt"

// for float multiplication

func multiply(no float64, n int) float64 {
	var ans = 1.0
	for i := 1; i <= n; i++ {
		ans = ans * no
	}
	return ans
}

func NthRoot(m, n float64) float64 {
	var low float64 = 1
	var high = m
	var mid float64
	// upper cap for the lowest value below which result an't go
	var eps float64 = 1e-6
	for (high - low) > eps {
		mid = low + (high-low)/2.0
		if multiply(mid, int(n)) < m {
			low = mid
		} else {
			high = mid
		}
		// fmt.Println(mid)
	}
	//fmt.Println(low)
	//fmt.Println(high)
	//fmt.Println(mid)
	return high
}

//	func multiply(no, n int) int {
//		var ans int = 1
//		for i := 1; i <= n; i++ {
//			ans = ans * no
//		}
//		return ans
//	}
//
//	func NthRoot(m, n int) int {
//		var low, high int = 1, m
//		var ans, mid int
//		for low <= high {
//			mid = low + (high-low)/2
//			if multiply(mid, n) <= m {
//				low = mid + 1
//				ans = mid
//			} else {
//				high = mid - 1
//			}
//		}
//		return ans
//	}
func reverse(revStr string, i, j int) string {
	var temp string = ""
	for k := i; k >= j; k-- {
		temp = tmep + revStr[k]
	}
	return temp
}
func main() {
	//var m = 16
	//var n = 2
	////ans := NthRoot(float64(m), float64(n))
	////fmt.Println(int(math.Ceil(ans)))
	//ans := NthRoot(float64(m), float64(n))
	//fmt.Println(ans)
	var s string = "Kiran"
	var rev []string
	fmt.Print(s)
	for i := 0; i < len(s); i++ {
		rev = append(rev, string(s[i]))
	}
	fmt.Print(rev)
	s = s + "k"
	fmt.Print(s)
}
