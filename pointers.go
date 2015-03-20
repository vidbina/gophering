package main

import "fmt"

func main() {
	var i int = 12
	p := &i

	fmt.Printf("p is a %T with the address ")
	fmt.Print(p)
	fmt.Printf(" containing ")
	fmt.Print(*p)
}
