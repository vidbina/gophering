package main

import "fmt"

func main() {
	var titanic Ship
	titanic = Boat{"Titanic"}
	fmt.Println(titanic.Dock())
	largeship := &titanic
	fmt.Println(Boat{"The Black Pearl"}.Dock())
	//fmt.Println(Aeroplane{"Air Force One"}.Dock())
}

// a few interfaces
type Honkable interface {
	Honk() string
}

type Driveable interface {
	Drive() string
}

type Flyable interface {
	LiftOff() string
}

type Submergable interface {
	Submerge() string
}

type Floatable interface {
	Dock() string
}

type Vehicle interface {
	Driveable
	Honkable
}

type Aeroplane interface {
	Driveable
	Flyable
}

type Ship interface {
	Floatable
	Honkable
}

type Bike interface {
	Driveable
	Honkable
}

type Train interface {
	Driveable
	Honkable
}

// a few structs
type Nameable struct {
	Name string
}

type Boat struct {
	Name string
}

func (b Boat) Dock() string {
	return "Docking the " + b.Name
}
func (b Boat) Honk() string {
	return "Hooonk Hooonk!!!"
}
