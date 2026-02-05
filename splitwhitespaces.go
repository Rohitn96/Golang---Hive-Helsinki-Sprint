package main

// import "fmt"

func SplitWhitespaces(s string) []string {
	var result []string
 	word := ""

 	for _, ch := range s {
 		if ch == ' ' || ch == '\t' || ch == '\n' {
 			if word != "" {
 				result = append(result, word)
 				word = ""
 			}
 		} else {
 			word += string(ch)
 		}
 	}

 	if word != "" {
 		result = append(result, word)
 	}

	return result
}

// func main() {
// 	fmt.Println(SplitWhitespaces("Hello! How have you been?"))
// }

