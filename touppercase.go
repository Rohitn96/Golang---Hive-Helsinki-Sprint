package main

// import "fmt"

func ToUpperCase(s string) string {
	r := []rune(s)
	result := ""
	for _,value := range r {
		if value >= 'a' && value <= 'z' {
			result += string(rune(value - 32))
		} else { result += string(value)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(ToUpperCase("Hello! How's your day going?"))
// }
