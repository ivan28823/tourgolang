package main

import "fmt"

type Vertext struct {
	Lat, Long float64
}

var m map[string]Vertext

var n = map[string]Vertext{
	"Bell Labs": Vertext{
		632446, 655165,
	},
	"Google": Vertext{
		54.65, 6516.54,
	},
}

var o = map[string]Vertext{
	"Bell Labs": {651651, 65165},
	"Google":    {5454.61, 116},
}

func main() {
	m = make(map[string]Vertext)
	m["Bell labs"] = Vertext{40.321, -46516.6341}
	fmt.Println(m["Bell labs"])

	// maps literals
	fmt.Println(n)

	// map literals continued
	fmt.Println(o)

	//mutating maps
	p := make(map[string]int)

	p["Answer"] = 42
	fmt.Println("The value", p["Answer"])

	p["Answer"] = 48
	fmt.Println("The value", p["Answer"])

	delete(p, "Answer")
	fmt.Println("The value", p["Answer"])

	v, ok := p["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
