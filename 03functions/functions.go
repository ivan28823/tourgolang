package main

import "fmt"

/*// simple function
func add(x int, y int) int {
	return x + y
}*/
//When two or more consecutive named function
//	parameters share a type, you can omit the
//	type from all but the last.
func add(x, y int) int {
	return x + y
}

// function with multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// main function
func main() {
	// print the output of the function
	fmt.Println("Add function output: ", add(7, 14))

	//swap function
	a, b := swap("world!", "Hello")
	fmt.Println("Swap function output: ", a, b)

	//split function
	fmt.Print("Split function output: ")
	fmt.Println(split(98))
}
