package sprint

// import "fmt"

func FilterBySum(arr [][]int, limit int) [][]int {

	result := [][]int {}
	for _,v := range arr{
		sum := 0
		for _, x := range v {
			sum += x
		}
	if sum >= limit {
			result = append(result, v)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9))
// }
