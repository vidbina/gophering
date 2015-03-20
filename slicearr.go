package main

import "fmt"

func main() {
	i := []string{"Mona", "An", "Gaby", "Titsia"}
	fmt.Printf("[]T %v is a %T\n", i, i)

	j := [...]string{"Mona", "An", "Gaby", "Titsia"}
	fmt.Printf("[...]T %v is a %T\n", j, j)

	k := [4]string{"Mona", "An", "Gaby", "Titsia"}
	fmt.Printf("[4]T %v is a %T\n", k, k)

	l := append(make([]string, 0), "Mona", "An", "Gaby", "Titsia")
	fmt.Printf("[]T %v is a %T\n", l, l)
}
