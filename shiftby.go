package sprint

func ShiftBy(r rune, step int) rune {
 bc := (r) + rune(step)

for bc > 122 {
   x := bc%122
   bc = 96 + x	
}
  return bc
}
