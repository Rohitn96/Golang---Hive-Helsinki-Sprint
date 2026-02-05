package main

// import "fmt"

func ToThePowerRecursive(n int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	result = (n * ToThePowerRecursive(n, power-1))
	return result
}

// func main() {
// 	fmt.Println(ToThePowerRecursive(2, 10))
// }
