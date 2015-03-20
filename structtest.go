package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
	}

	//fmt.Print(Point{12, 33} + Point{8, 7})

	p := Point{12, 5}
	ptr := &p
	fmt.Printf("point x is a %T containing %v, ptr to x is a %T containing %v", p, p.X, ptr, ptr.X)
}
