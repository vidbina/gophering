package main

import "fmt"

type Character struct {
	Name         string
	Interjection string
}

func (c Character) Greet() string {
	return "Hello I'm " + c.Name
}

func (c Character) Interject() string {
	return c.Interjection + "!"
}

func main() {
	Cast := []Character{
		Character{Name: "Homer", Interjection: "Doh"},
		Character{Name: "Moe", Interjection: "Wh-wh-whaat"},
		Character{Name: "Monty Burns", Interjection: "Excellent"},
	}
	stewie := &Character{"Steward", "Blast"}
	fmt.Println(Cast[1].Greet())
	fmt.Println(stewie.Interject())
}
