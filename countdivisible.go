package sprint

// import "fmt"

func CountDivisible(from, to, step, divisor int) int {
	if step <= 0 || divisor == 0 {
		return 0
	}
	var count int
 	for i := from ; i < to ; i += step {
          if i%divisor = 0{
          count ++
          }
  }
  return count
}

// func main() {
// 	fmt.Println(CountDivisible(5, 17, 2, 3))
// }
