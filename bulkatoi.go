package main

// import "fmt"

func BulkAtoi(arr []string) []int {
	var out []int
	for _, s := range arr {
		out = append(out, StrToInt(s))

	}
	return out
}

func StrToInt(s string) int {

	// checking for empty string
	if len(s) == 0 {
		return 0
	}
	sign := 1
	index := 0

	// Handle optional signs

	if s[0] == '-' {
		sign = -1
		index++
	} else if s[0] == '+' {
		index++
	}

	// Reject multiple signs
	if index < len(s) && (s[index] == '+' || s[index]=='-') {
		return 0
	}
	result := 0

	// Convert Characters to digits
	for index < len(s) {
		c:= s[index]

		if c < '0' || c > '9' {
			return 0
		}
		digit := int(c - '0')
		result = result*10 + digit

		index++
	}

	return sign * result
}

// func main() {
// 	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
// }
