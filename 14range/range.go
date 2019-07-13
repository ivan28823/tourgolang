package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// range continued

	pw := make([]int, 11)
	// only the index
	for i := range pw {
		pw[i] = 1 << uint(i)
	}
	// skip index
	for _, value := range pw {
		fmt.Println(value)
	}
}
