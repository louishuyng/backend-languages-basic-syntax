package test

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Methods() {
	fmt.Println("Test Methods")
	v := Vertex{3, 4}
	fmt.Printf("  v.Abs()=%v\n", v.Abs())

	fmt.Println("Test Methods Are Just a Function")

	abs := func(v Vertex) float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	fmt.Printf("  abs(v)=%v\n", abs(v))

	fmt.Println("Test Methods on non-struct type")

	f := MyFloat(-math.Sqrt2)
	fmt.Printf("  f.Abs()=%v\n", f.Abs())
}
