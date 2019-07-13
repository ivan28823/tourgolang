package main

import (
	"fmt"
	"time"
)

func main() {
	//simple for statement
	for i := 0; i < 3; i++ {
		fmt.Println("In for loop for ", i+1, "times")
	}

	//another for loop, not necesary init statement
	// for is go's "while"
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("Sum: ", sum)

	//run forever
	l := 0
	for {
		fmt.Println("infinete loop")
		time.Sleep(50 * time.Millisecond)
		l++
		if l > 10 {
			break // break the loop
		}
	}
}
