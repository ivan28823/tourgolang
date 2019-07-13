package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]

	fmt.Println("print slice ", s)

	// slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "changed"
	fmt.Println(a, b)
	fmt.Println(names)

	//slice literals

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q:", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("r:", r)

	sl := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("sl:", sl)

	//slice default
	o := []int{2, 3, 5, 7, 11, 13}

	o = o[1:4]
	fmt.Println("o = o[1:4]>>", o)

	o = o[:2]
	fmt.Println("o = o[:2]>>", o)

	o = o[1:]
	fmt.Println("o = o[1:]>>", o)

	// slice length
	ot := []int{2, 3, 5, 7, 11, 13}
	printSlice(ot)

	// slice the slice to give it zero length
	ot = ot[:0]
	printSlice(ot)

	//extends its length
	ot = ot[:4]
	printSlice(ot)

	//drop its first two values
	ot = ot[2:]
	printSlice(ot)

	//nil slices
	var nl []int
	printSlice(nl)
	if nl == nil {
		fmt.Println("nil slice")
	}

	// creating slice with make
	z1 := make([]int, 5)
	printSlice2("z1", z1)

	z2 := make([]int, 0, 5)
	printSlice2("z2", z2)

	z3 := z2[:2]
	printSlice2("z3", z3)

	z4 := z3[2:5]
	printSlice2("z4", z4)

	// slice of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for t := 0; t < len(board); t++ {
		fmt.Printf("%s\n", strings.Join(board[t], " "))
	}

	// appending to a slice
	var ap []int
	printSlice(ap)

	//append works on nil slice
	ap = append(ap, 0)
	printSlice(ap)

	// the slice grows as needed
	ap = append(ap, 1)
	printSlice(ap)

	//we can add more than one element at a time
	ap = append(ap, 2, 3, 4)
	printSlice(ap)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
