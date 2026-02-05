package main

// import "fmt"

func FactorialIterative(n int) int {
	if n <= 1 {
		return 1
	}
	result := 0
	for i := n ; i > 1; i-- {
	result = result*i
	}
	return result
}

// func main() {
// 	fmt.Println(FactorialIterative(5))
// }
