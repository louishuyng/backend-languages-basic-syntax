package test

import (
	"fmt"
	"runtime"
	"time"
)

func Switch() {
	fmt.Println("Test Switch")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("  OS X.")
	case "linux":
		fmt.Println("  Linux.")
	default:
		fmt.Printf("  %s.\n", os)
	}
	fmt.Println("Test Switch evaluation order")
	fmt.Println("  When's Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("  Today.")
	case today + 1:
		fmt.Println("  Tomorrow.")
	case today + 2:
		fmt.Println("  In two days.")
	default:
		fmt.Println("  Too far away.")
	}

	fmt.Println("Test Switch no condition")

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("  Good morning!.")
	case t.Hour() < 17:
		fmt.Println("  Good afternoon!.")
	default:
		fmt.Println("  Good evening!.")
	}
}
