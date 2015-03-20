package main

import "fmt"

type Employee struct {
	Name    string
	Company Company
	Role    string
}

type Company struct {
	Name string
	Type string
}

func main() {
	saddlco := Company{"Saddl", "BV"}
	saddlteam := map[string]Employee{
		"Max":    Employee{"Max", saddlco, "Product"},
		"George": Employee{"George", saddlco, "CEO"},
		"Linda":  Employee{"Linda", saddlco, "Marketing"},
		"Orgeon": Employee{"Orgeon", saddlco, "Business Development"},
		"Eleni":  Employee{"Eleni", saddlco, "Front-end Development"},
	}
	fmt.Print(saddlteam)
}
