package main

import "fmt"


func GenerateRange(min, max int) []int {
	if min >= max { 
		return nil
	}

//	length := max - min

	newslice := make([]int, 0)

	for i := min ; i < max ; i++ {
		newslice = append(newslice, i)
	}
	return newslice
}

func main() {
	fmt.Println(GenerateRange(4, -1))
}
