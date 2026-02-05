package main

import "fmt"

func BalanceOut(arr []bool) []bool {
	countTrue := 0
	countFalse := 0

	for _,value := range arr{
		if value { countTrue++
		} else {
			countFalse++
		}
	}
	
	if countTrue > countFalse {
		diff := countTrue - countFalse
		for i:=0 ; i<diff ;i++{
			arr = append(arr, false)
		}
	}

	if countFalse > countTrue {
		diff := countFalse - countTrue
		for i:=0 ; i<diff ; i++{
			arr = append(arr, true)
		}
	}
	return arr
}

func main() {
	fmt.Println(BalanceOut([]bool{true, false, false, false}))
}
