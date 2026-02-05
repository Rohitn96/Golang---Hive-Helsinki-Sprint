package main

// import "fmt"

func StrConcatWith(strs []string, sep string) string {
	abc := ""

	for i, value := range strs {
		if i < len(strs)-1 {
			value += sep
			}
			abc += value
		}
	return abc


}

// func main() {
// 	fmt.Println(toConcat := []string{"Three", " Two", " One", " Go!"})
// }
