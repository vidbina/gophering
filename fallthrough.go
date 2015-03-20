package main

import (
	"fmt"
)

func main() {
	res := "something"
	switch res {
	case "something":
		fmt.Print("found something")
		fallthrough
	case "else":
		fmt.Print("something else")
	default:
		fmt.Print("unknown")
	}
}
