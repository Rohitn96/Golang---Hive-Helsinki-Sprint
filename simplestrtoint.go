package main

// import "fmt"

func SimpleStrToInt(s string) int {
	r := 0

	for _, ch := range s{
		if ch < '0' || ch > '9' {
			return 0
		}
		r = r*10 + int(ch - '0')
	}
	return r
}

// func main() {
// 	fmt.Println(SimpleStrToInt("00000010533"))
// }
