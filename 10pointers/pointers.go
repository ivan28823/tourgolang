package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println("Access to i via p", *p, p)
	*p = 21
	fmt.Println("changed i via p", i)

	p = &j
	*p = *p / 37
	fmt.Println("changed j via p", j)
}
