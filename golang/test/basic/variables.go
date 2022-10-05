package test

import "fmt"

var c, python, java bool

func Variables() {
	var i, j = 1, 2

	fmt.Printf("Test Variables c=%v python=%v java=%v i=%v j=%v \n", c, python, java, i, j)
}
