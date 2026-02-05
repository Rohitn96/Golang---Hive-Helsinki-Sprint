package main

// import "fmt"

func FactorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// func main() {
// 	fmt.Println(FactorialRecursive(10))
// }
