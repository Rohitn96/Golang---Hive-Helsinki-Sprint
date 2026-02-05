package main

import "fmt"

func main() {
	fmt.Println(ToCapitalCase("Hello! GreAt to sEe you! HOw-aRe-you-doing-2day?"))
}

func ToCapitalCase(s string) string {
	result := []rune{}
	ab := true
	 //  The same loop will go through and check if there is non alphanumeric characters
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			// Once we have it we make the changes to upper and lower case.
			if ab {
				if v >= 'a' && v <= 'z' {
					v = v - 32
				}
				ab = false
			} else {
				if v >= 'A' && v <= 'Z' {
					v = v + 32
				}
			}
		} else {
			ab = true
		}
		//fmt.Println(string(v))
		result = append(result, v)
	}

	return string(result)
}

// func main() {
// 	fmt.Println(StrReverse("Hello Coder!"))
// }
