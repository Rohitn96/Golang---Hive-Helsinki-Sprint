
package sprint

func ReverseAlphabet(step int) string {
    if step <= 0 {
    step = 1
 }

  var rev string
    for i:= 'z' ; i >= 'a' ; i = i - rune(step) {
      rev = rev + string(rune(i))
}
return  rev
}

