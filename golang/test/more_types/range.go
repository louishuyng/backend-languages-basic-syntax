package test

import "fmt"

func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 18}

	fmt.Println("Test Range")
	for i, v := range pow {
		fmt.Printf("  2**%d = %d\n", i, v)
	}

	fmt.Println("Test Range continued")
	pow = make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("  %d\n", value)
	}
}
