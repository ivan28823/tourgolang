package main

import "fmt"

func fibbonacci(n int, c chan int64) {
	var x, y int64 = 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int64, 50)
	go fibbonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
