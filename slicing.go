package main

import "fmt"

func main() {
	a := [5]int{12, 32, 15, 74, 1}
	fmt.Println("bounded ", a)

	b := []int{12, 32, 15, 74, 1}
	fmt.Println("unbounded", b)

	c := make([]int, 5)
	fmt.Println("unbounded", c)

	d := make([]int, 5, 10)
	fmt.Println("unbounded", d)
}
