package main

// import "fmt"

func ToThePowerIterative(n int, power int) int {

	result := 1

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	for i := power ; i > 0 ; i-- {
		result *= n
	}
	return result
}

// func main() {
// 	fmt.Println(ToThePowerIterative(2, 10))
// }
