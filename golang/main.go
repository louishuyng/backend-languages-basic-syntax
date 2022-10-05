package main

import (
	basic "BackendGolang/test/basic"
	flow_controls "BackendGolang/test/flow_controls"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the playgrond!")

	basic.Packages()
	basic.Functions()
	basic.NameReturnValues()
	basic.Variables()
	basic.ShortVariableDeclarations()
	basic.BasicTypes()
	basic.ZeroValues()
	basic.TypeConversions()
	basic.Constants()
	basic.Constants()
	basic.NumericConstants()

	flow_controls.For()
	flow_controls.If()
	flow_controls.Switch()
	flow_controls.Defer()
	flow_controls.StackingDefer()
}
