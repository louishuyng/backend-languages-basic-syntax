package test

import "fmt"

func Defer() {
	defer fmt.Print(" world\n")

	fmt.Println("Test Defer")
	fmt.Print("  Hello")
}
