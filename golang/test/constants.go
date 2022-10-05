package test

import "fmt"

const Pi = 3.14

func Constants() {
	const World = "Hello world"
	fmt.Println("Test Constants")

	fmt.Println("  Hello", World)
	fmt.Println("  Happy", Pi, "Day")

	const Truth = true
	fmt.Println("  Go rules?", Truth)
}
