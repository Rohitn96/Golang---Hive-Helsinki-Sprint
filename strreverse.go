package sprint

// import "fmt"

func StrReverse(s string) string {
	r := []rune(s)

// the below line takes 2 elements in the for loop,
// it runs in reverse order, until it reaches a middle point and then we swap.
	for i,j := 0, len(s)-1 ; i < j ; i,j = i+1, j-1 {
		// fmt.Println("%v ," , i, j )
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// func main() {
// 	fmt.Println(StrReverse("Hello Coder!"))
// }
