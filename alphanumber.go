package main

import "fmt"

func AlphaNumber(n int) string {
	abc := ""
	if n == 0 {
		return "a"
	} 
	if n < 0 {
		abc += "-"
		n = n* -1
	}
	result := ""
	for n > 0 {
		digit := n%10
		char := string(rune('a'+ digit))
		result = char+ result
		n = n/10 
	}
	return abc + result
}

func main() {
	fmt.Println(AlphaNumber(0))
}