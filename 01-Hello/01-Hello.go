package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	var t int
	t = 0
	fmt.Println(&t)
	for i := 0; i < 10; i++ {
		t = i
		fmt.Println("variant out for-loop address:", &t)

	}
	for i := 0; i < 10; i++ {
		p := i
		fmt.Println("variant in for-loop address:", &p)
	}
}
