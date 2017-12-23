package main

import (
	"fmt"
)

func convert() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}
func main() {
	a := true
	b := false
	fmt.Println("ä:", a)
	c := a && b
	fmt.Println("c", c)
	x := 1 + 1i
	y := -1 - 1i
	fmt.Println(x, x+y, x*y)

	const trueConst = true
	type myBool bool

	var customTrue myBool = trueConst
	fmt.Println(customTrue)
	fmt.Printf("a's type %T", a)
	convert()
}
