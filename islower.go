package sprint

// import "fmt"

func IsLower(s string) bool {
	r := []rune(s)
	ab := true
	for _, value := range r{
		if value < 'a' || value > 'z' {
			ab = false
	}
	}
	return ab
}

// func main() {
// 	fmt.Println(IsLower("hello"))
// 	fmt.Println(IsLower("world!"))
// }
