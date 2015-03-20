package main

import "fmt"

func main() {
	var x string = "before"
	defer fmt.Println("testing ", x)
	for i := 0; i < 5; i++ {
		defer fmt.Println("currently ", i)
	}
	x = "after"
	fmt.Println("testing ", x)
}
