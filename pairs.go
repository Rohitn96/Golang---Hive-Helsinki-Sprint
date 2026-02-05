package main
import "fmt"

func Pairs() string {
    var pa string
    for i := 0 ; i < 98 ; i++ {
        for j := i + 1 ; j <= 99 ; j++ {
                pa += fmt.Sprintf("%02d %02d, ", i, j)
        }
    }
    return pa + ("98 99")
}
func main() {
  fmt.Println(Pairs())
}
