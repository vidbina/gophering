package main

import "fmt"

type Car struct {
	Brand, Model, Type string
	Wheels             int
}

func (car *Car) Rebrand(brand string) {
	car.Brand = brand
}

func (car Car) Rebase(count int) {
	car.Wheels = count
}

func main() {
	fmt.Println(Car{})
	tesla := Car{Brand: "Tsla", Model: "X", Type: "electric", Wheels: 4}
	tesla.Rebrand("Tesla")
	fmt.Println(tesla)

	mack := Car{Brand: "Mack", Model: "Semi", Type: "diesel", Wheels: 4}
	mack.Rebase(16)
	fmt.Println(mack)
}
