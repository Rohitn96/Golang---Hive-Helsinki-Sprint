package sprint

func ReverseAlphabetValue(ch rune) rune {
   const a = 'a'
   const z = 'z' 
   revers := int(z) - int(ch-a)
   return rune(revers)
}
