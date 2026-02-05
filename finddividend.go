package sprint

func FindDividend(from, to, divisor int) int {

	for i := from ; i < to ; i++ {
		if i%divisor == 0 {
			return i
		}
	}
	return -1
}


// func main() {
// 	fmt.Println(FindDividend(10, 18, 9))
// }
