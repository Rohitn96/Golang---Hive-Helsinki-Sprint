package main

import "fmt"

func Combinations() string {
	var comb string
	for i := 0 ; i <  ; i++ {
		for j := i + 1 ; j < 9 ; j++ {
			for k := j + 1 ; k < 10 ; k++ {

				comb += fmt.Sprintf("%d%d%d, ", i, j, k)

			}
		}
	}
	comb = comb[:len(comb)-2]
    return comb
}

func main() {
	fmt.Println(Combinations())
}
