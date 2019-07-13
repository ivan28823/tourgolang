package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods are functions
func Abs_f(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods continued
type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//pointer recivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//function and pointers
func Scale_f(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("print method", v.Abs())

	//methods are functions
	fmt.Println("print method", Abs_f(v))

	//methods continued
	f := Myfloat(-math.Sqrt2)

	fmt.Println("methods continued", f.Abs())

	//pointer reciver
	v.Scale(10)
	fmt.Println("ponter reciver", v.Abs())

	//pointers and functions
	u := Vertex{3, 4}
	Scale_f(&u, 10)
	fmt.Println("pointers and function", Abs_f(u))

	//methods and pointer indirection
	p := &Vertex{3, 4}
	p.Scale(5)
	Scale_f(p, 2)
	fmt.Println(u, p)

}
