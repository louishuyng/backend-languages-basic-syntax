package test

import "fmt"

func StackingDefer() {
	fmt.Println("Test Stacking Defer")

	fmt.Println("  counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("  i=%v", i)
	}

	fmt.Println("  done")
}
