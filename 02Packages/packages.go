//package name, go run the main package
package main

// import used packages
import (
	"fmt"
	"math"
	"math/rand"
)

// this code also is valid
//import "fmt"
//import "math/rand"

func main() {
	// using imported packages
	fmt.Println("Random number: ", rand.Intn(100))

	// use an exported name, starts with capital letter
	fmt.Println("Pi is", math.Pi)
}
