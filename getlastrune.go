package sprint

// import "fmt"

func GetLastRune(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-1]

}

// func main() {
// 	fmt.Printf("%c\n", GetLastRune("kood"))
// }
