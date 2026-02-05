package main

// import "fmt"

func IsNumeric(s string) bool {
	r := []rune(s)
	ab := true

	for _, value := range r{
		if value < '0' || value > '9' {
			return false
		}
	}
	return ab
}


// func main() {
// 	fmt.Println(IsNumeric("012345"))
// 	fmt.Println(IsNumeric("01+23+45"))
// }
