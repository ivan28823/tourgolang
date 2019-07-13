package main

/**	Basic types
 *	bool
 *
 *	string
 *
 *	int  int8  int16  int32  int64
 *	uint uint8 uint16 uint32 uint64 uintptr
 *
 *	byte // alias for uint8
 *
 *	rune // alias for int32
 *	     // represents a Unicode code point
 *
 *	float32 float64
 *
 *	complex64 complex128
 */
import (
	"fmt"
	"math"
	"math/cmplx"
)

// declaration
var c, python, java bool

// declaration with initializer
var k, l int = 1, 2

// diferent types
var str1, num, bolean = "string", 2.5, false

// several types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	var i int
	// short variable declaration only inside a function
	shortvar := "short declaration"
	fmt.Println(i, c, python, java, k, l, str1, num, bolean, shortvar)

	//print variables type
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)

	// zero values
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Printf("Zero values: %v %v %v %q\n", zi, zf, zb, zs)

	//type conversions
	var tx, ty int = 3, 4
	var tf float64 = math.Sqrt(float64(tx*tx + ty*ty))
	var tz uint = uint(tf)
	fmt.Println(tx, ty, tz)

	//Type interface
	tivar := "string"
	fmt.Printf("tivar is of type %T\n", tivar)

	// constants
	const Pi = 3.14
	const World = "World"
	fmt.Println("hello", World)
	fmt.Println("Pi value", Pi)

	// big constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
