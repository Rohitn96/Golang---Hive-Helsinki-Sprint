package main

import "fmt"

func AlphabetMastery(n int) string {
   var alp string
   for i := 0 ; i <= n ; i++ {
     alp = alp + string(('a') + rune(i))
  }
 return alp
}

func main() { fmt.Println(AlphabetMastery(6)) }
