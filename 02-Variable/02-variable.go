package main

import (
	"math"
	"fmt"
)
func main() {
	//1 declaration
	var age int

	fmt.Println("age is", age)
	//2 assign
	age = 18
	fmt.Println("age is", age)
	var name string = "jared"
	fmt.Println("name is ", name)

	var width, height int = 50, 100
	fmt.Printf("width = %d, height = %d \n", width, height)


	var (
		sex = "formal"
		weight = 60
	)

	fmt.Println("sex", sex, "weight", weight)

	a, b := 1.00000000001, 1.000000001
	fmt.Println("a, b", a, b)
	c:= math.Min(a, b)
	fmt.Println(c)
}