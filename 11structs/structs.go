package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1  = Vertex{1, 2}
	v2  = Vertex{X: 1}
	v3  = Vertex{}
	ptr = &Vertex{1, 2}
)

func main() {
	// simple
	fmt.Println(Vertex{1, 2})

	// struct fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("access via dot", v.X)

	//pointers to struct
	p := &v
	p.X = 1e9
	fmt.Println("changed values via pointer", v)

	// struct literals
	fmt.Println("struct literals", v1, v2, v3, ptr)
}
