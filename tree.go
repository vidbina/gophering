package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
	//	"time"
)

//type Tree struct {
//	Left  *Tree
//	Value Int
//	Right *Tree
//}

func Walk(t *tree.Tree, ch chan int, level int) {
	fmt.Printf("\n%3d", level)
	for i := 0; i < level; i++ {
		fmt.Printf("..")
	}

	// always walk left branch and report values until nil
	left, right := make(chan int), make(chan int)

	if t != nil {
		fmt.Printf(" node %d:[%v]", t.Value, t)
		ch <- t.Value

		go Walk(t.Left, left, level+1)
		go Walk(t.Right, right, level+1)

		complete := map[string]bool{
			"left":  false,
			"right": false,
		}

		for {
			if complete["left"] && complete["right"] {
				close(ch)
				return
			}

			select {
			case node, ok := <-left:
				if ok {
					ch <- node
				} else {
					complete["left"] = true
				}
			case node, ok := <-right:
				if ok {
					ch <- node
				} else {
					complete["right"] = true
				}
			default:
			}
		}
	} else {
		// close my channel if node is nil
		close(ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	t := tree.New(1)

	t1, t2 := make(chan int), make(chan int)
	go Walk(t, t1, 0)

	one, two := make(map[int]bool), make(map[int]bool)

	for {
		select {
		case val, ok := <-t1:
			if !ok {
				return
			}
			one[val] = true
			fmt.Printf(" got(%v %v)", val, ok)
		case val, ok := <-t2:
			if !ok {
				return
			}
			two[val] = true
			fmt.Printf(" got(%v %v)", val, ok)
		default:
		}
	}

	fmt.Printf("1: %v", one)
	fmt.Printf("2: %v", two)
	/*
		for {
			select {
			case val := <-t1:
				one[val] = true
				fmt.Println(one)
			case val := <-t1:
				two[val] = true
				fmt.Println(two)
			default:
				time.Sleep(1 * time.Second)
			}
		}
	*/
}
