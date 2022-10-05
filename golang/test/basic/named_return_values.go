package test

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func NameReturnValues() {
	fmt.Print("Test named return values split 7 => ")
	fmt.Println(split(7))
}
