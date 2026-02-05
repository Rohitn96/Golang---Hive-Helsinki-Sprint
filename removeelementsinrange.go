package main

import "fmt"

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	if to < from {
		to, from = from, to
	}

	result := make([]float64, 0)
	result = append(arr[:from], arr[to:]...)
	// for i := 0 ; i < len(arr) ; i++ {
	// 	if i < from || i >= to {
	// 		result = append(result, arr[i])
	// 	}
	// }
		return result
}

	func main() {
		fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
	}
