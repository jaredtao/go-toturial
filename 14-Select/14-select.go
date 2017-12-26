package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server 1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}
func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	go server1(out1)
	go server2(out2)
	select {
	case s1 := <-out1:
		fmt.Println(s1)
	case s2 := <-out2:
		fmt.Println(s2)
	}
}
