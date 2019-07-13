package main

import "fmt"

func fibonacci() func() int {
	index := -1
	n1, n2 := 0, 1
	return func() int {
		index++
		if index <= 1 {
			return index
		} else {
			n2, n1 = n2+n1, n2
			return n2
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
