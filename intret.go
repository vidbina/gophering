package main

import "fmt"

type Starlike interface {
	Glow() string
	Fall() string
}

type Star struct {
	Name  string
	Color string
}

func main() {
	var glower Starlike
	glower = &Star{Name: "L32", Color: "blue"}
	fmt.Println("The star we got is", glower)
	fmt.Println(glower.Glow())
	fmt.Println(glower.Fall())
}

func (s *Star) Glow() string {
	return "shine"
}

func (s *Star) Fall() string {
	return "pixie dust"
}
