package main

import "fmt"

func finished() {
	fmt.Println("finsihed")
}

func largest(nums []int) {
	defer finished()

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("max", max)
}

func deferStack() {
	name := "Naveen"
	fmt.Println("Origin string ", name)
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
func main() {
	nums := []int{1, 23123, 3213, 312, 321}
	largest(nums)
	deferStack()
}
