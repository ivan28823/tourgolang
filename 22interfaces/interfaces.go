package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertext struct {
	X, Y float64
}

func (v *Vertext) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// interfaces are implemented implicitly
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	fmt.Println(t.S)
}

// interface value
/*func (t *T) M2() {
	fmt.Println(t.S)
}*/

type F float64

func (f F) M() {
	fmt.Println(f)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertext{3, 4}

	a = f  // a My float implements Abser
	a = &v //a *Vertext implenets Abser

	// In the following line, v is a Vertext (not *Vertext)
	// and does NOT implement Abser
	//a = &v

	fmt.Println(a.Abs())

	// implicitly
	/*var i I = T{"hello"}
	i.M()*/

	// inteface values
	var j I

	j = &T{"Hello"}
	describe(j)
	j.M()

	j = F(math.Pi)
	describe(j)
	j.M()
}
