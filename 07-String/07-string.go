package main

import "fmt"

func mutate(s []rune) string {
	s[0] = 's'
	return string(s)
}
func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
