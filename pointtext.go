package main

import "fmt"

type Point struct{
    X float32
    Y float32
    Text string
}


func PointText(p Point) Point {
	p.Text = fmt.Sprintf("Text at X=%f, Y=%f)", p.X, p.Y)
	return p
}

// func main() {
// 	var P Point
// //	var Q Point
// //	var Z Point

// 	P.X = 1
// 	P.Y = 2


// 	fmt.Println(PointText(P))
// }
