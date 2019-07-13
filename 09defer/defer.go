package main

import "fmt"

func main() {
	//A defer statement defers the execution
	//	of a function until the surrounding function returns.
	defer fmt.Println("world")

	fmt.Print("Hello ")

	fmt.Println("Counting...")
	//Deferred function calls are pushed onto a stack.
	//	When a function returns, its deferred calls are
	//	executed in last-in-first-out order.
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}
