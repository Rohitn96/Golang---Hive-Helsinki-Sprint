package sprint 

// import "fmt"

func BetweenLimits(from, to rune) string {
     start := from
     end := to
     if from > to {
     start, end = end, start }
       output := "" 
    for i := start + 1 ; i < end ; i++ {
        output = output + string(i)
 }
     return output
}

func main() {
     fmt.Println(BetweenLimits('z','a'))
}
