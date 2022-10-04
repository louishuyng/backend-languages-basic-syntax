package test

import (
	"fmt"
	"math"
)

func TypeConversions() {
	x, y := 3, 4

	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)

	fmt.Println("Test Type conversions")
	fmt.Printf("  %v %v %v", x, y, z)
}
