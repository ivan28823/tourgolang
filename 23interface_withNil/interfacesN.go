package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	//empty interface
	fmt.Println("Print empty interface")
	var j interface{}
	desc(j)

	j = 42
	desc(j)

	j = "hello"
	desc(j)
}

func desc(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
