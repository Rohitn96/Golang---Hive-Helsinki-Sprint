package sprint


func Accumulate(n int) int {
 if n < 0 {return 0}
  acc := 0
 for i := 0 ; i <= n ; i++ {
    acc  += i
}
    return acc
}

