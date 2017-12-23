package main

import (
	"fmt"
)

func useSlice() {
	a := []int{12, 2342, 3242, 242, 342, 423}
	var b []int = a[1:4]
	fmt.Println(b)
}
func mulitDimensional() {
	s := [][]int{{1, 2, 3}, {2, 3, 4}, {10, 12, 1, 4}}
	for _, v1 := range s {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
func useCap() {
	fruitarray := []string{"apple", "orange", "banana", "soup", "belo", "qwe"}
	fruitslice := fruitarray[1:3]
	fmt.Println(fruitarray, len(fruitarray))
	fmt.Println(fruitslice, len(fruitslice), cap(fruitslice))
}
func useMake() {
	i := make([]int, 4, 5)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
	i = append(i, 1)
	fmt.Println(i, len(i), cap(i))
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 1
	}
}
func useCopy() {
	countries := []string{"USA", "CHINA", "JAPAN", "KOERA"}
	fmt.Println("countries has length:", len(countries), "and cap", cap(countries))
	needCountries := countries[:len(countries)-2]
	fmt.Println("needCountries has length:", len(needCountries), "and cap", cap(needCountries))
	needCountries = append(needCountries, "UK")
	needCountries = append(needCountries, "UK")
	needCountries = append(needCountries, "UK")
	fmt.Println("needCountries has length:", len(needCountries), "and cap", cap(needCountries))

	countriesCopy := make([]string, len(needCountries))
	copy(countriesCopy, needCountries)
	fmt.Println("countriesCopy has length:", len(countriesCopy), "and cap", cap(countriesCopy))
	fmt.Println(countriesCopy)
	fmt.Println("countries has length:", len(countries), "and cap", cap(countries))
}
func main() {
	var a [3]int
	for i, v := range a {
		fmt.Println(i, v)
	}
	b := [3]int{12, 78, 50}
	fmt.Println(b)

	c := [3]int{12}
	fmt.Println(c)

	mulitDimensional()
	useSlice()
	useCap()
	useMake()

	d := []int{1, 2, 3, 4}
	fmt.Println(d, len(d))
	subtactOne(d)
	fmt.Println(d, len(d))

	e := d[1:3]
	fmt.Println("e", e, len(e))
	subtactOne(e)
	fmt.Println("e", e, len(e))

	useCopy()
}
