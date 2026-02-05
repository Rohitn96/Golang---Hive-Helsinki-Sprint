package sprint

// import "fmt"

func StrLength(s string) []int {

	runes := []rune(s)

	xyz := len(s)

	ab:= []int{len(runes),xyz}
	return ab

}


// func main() {
// 	fmt.Println(StrLength("Hello 😁 World!"))
// }
