package test

import "fmt"

func For() {
	fmt.Println("Test For")
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("  sum=%v\n", sum)

	fmt.Println("Test For continued")

	sum = 1

	for sum < 100 {
		sum += sum
	}

	fmt.Printf("  sum=%v\n", sum)
}
