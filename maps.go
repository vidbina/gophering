package main

import (
	"fmt"
)

type Person struct {
	Name, City  string
	Age, Weight int
}

//var simpsons
func main() {
	//simpsons := make(map[string]Person)
	//simpsons["Homer"] = Person{Name: "Homer J Simpson"}
	//simpsons["Bart"] = Person{Name: "Bartolomew Simpson"}

	//simpsons := map[string]Person{
	//	"Homer": Person{Name: "Homer J Simpson"},
	//	"Bart":  Person{Name: "Bartolomew Simpson"},
	//}

	simpsons := map[string]Person{
		"Homer": {Name: "Homer J Simpson"},
		"Bart":  {Name: "Bartolomew Simpson"},
	}

	fmt.Print(simpsons)
	lisa, ok := simpsons["Lisa"]
	fmt.Print("Lisa exists ", ok)
}
