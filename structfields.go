package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Alive bool
}

func main() {
	stewie := Person{"Steward Gilligan Griffin", 2, true}
	fmt.Print(stewie)

	// NOTE: impossible to mix field:value and value initializers
	//peter := Person{Name: "Peter Griffin", 46, true}
	lois := Person{Name: "Lois Griffin", Age: 42, Alive: true}
	fmt.Print(lois)
}
