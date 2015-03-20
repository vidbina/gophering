package main

import "fmt"

type Character struct {
	Name, Series string
}

func (c Character) Introduce() string {
	return ("Hello my name is " + c.Name)
}

//func (c *Character) Introduce() string {
//return ("I am " + c.Name + " and you know me from " + c.Series)
//}

func main() {
	fmt.Println(Character{"Peter Griffin", "Familyguy"}.Introduce())
	fmt.Println(Character{"Bart Simpson", "The Simpsons"}.Introduce())
	fmt.Println(Character{"Roger Smit", "American Dad"}.Introduce())
	fmt.Println(Character{"Leopold Butters Stotch", "Southpark"}.Introduce())

	c := Character{"Rollo", "Stoolbend"}
	p := &c
	fmt.Println(p.Introduce())
}
