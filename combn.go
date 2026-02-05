package main

import "fmt"


func CombN(n int) []string {
	if n < 1 || n > 10 {
		return nil
	}

	var res []string
	var backtrack func(int, string)

	backtrack  = func(start int, current string) {
		if len(current) == n {
			res = append(res, current)
			return
		}
		for i := start ; i <= 9 ; i++ {
			backtrack(i+1, current+string(rune('0'+i)))
		}
	}
	backtrack(0, "")
	return res
}

func main() {
	fmt.Println(CombN(4))
}

	// 	var comb string
// 	for i := 0 ; i < 1 ; i++ {
// 		for j := i + 1 ; j < 2 ; j++ {
// 			for k := j + 1 ; k < 3 ; k++ {
// 				for l := k + 1 ; l < 4 ; l++ {
// 					for m := l + 1 ; m < 5 ; m++ {
// 						for n := m + 1 ; m < 6 ; n++ {
// 							for o := n + 1 ; o < 7 ; o++ {
// 								for p := o + 1 ; p < 8 ; p++ {
// 									for q := p + 1 ; q < 9 ; q++ {
// 										comb += fmt.Sprintf("%d%d%d%d%d%d%d%d%d, ", i, j, k, l, m, n, o, p, q)
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	comb = comb[:len(comb)-2]
//     return string(comb)
// }
