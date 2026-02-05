package main

// import "fmt"

func StrCompare(a, b string) int {
	if a == b {
		return  0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

// func main() {
// 	fmt.Println(StrCompare("Day", "ay"))
// }
