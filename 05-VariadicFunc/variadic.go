package main

import "fmt"

func find(num int, nums ...interface{}) {
	fmt.Printf("type of num is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "fount at index ", i, "in", nums)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("not found")
	}
	fmt.Println()
}
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}
func main() {
	find(89, 1, 12, 2)
	find(1, 2, 3, 1, 4)
	find(1)
	find(1, 2.4, "123")

	welcome := []string{"Hello", "World"}
	change(welcome...)
	fmt.Println(welcome)
}
