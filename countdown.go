package main

import "fmt"

func Countdown(n int) string {
	var cd string
	for i := n ; i >= 1 ; i -= 2 {
		cd += string('0' + rune(i)) + ", "
	}
	return cd + "0!"
}
func main() {
	fmt.Println(Countdown(9))
}
