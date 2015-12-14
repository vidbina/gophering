package main

import "fmt"

type Car struct {
	Wheels, Doors int
	Roof          bool
	Fuel          string
}

func (c Car) Description() Car {
	fmt.Printf("The %v-wheeled %v-door %q car", c.Wheels, c.Doors, c.Fuel)
	return c
}

func main() {
	rob := Car{4, 5, true, "Diesel"}
	fmt.Printf("\nRoof? %v", rob.Description().Roof)
}
