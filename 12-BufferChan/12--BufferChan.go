package main

import (
	"fmt"
	"sync"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("succeed write ", i, " to chan", len(ch), cap(ch))
	}
	close(ch)
}
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
func main() {
	{
		//		ch := make(chan int, 3)
		//		go write(ch)
		//		time.Sleep(2 * time.Second)
		//		for v := range ch {
		//			fmt.Println("read v", v)
		//			time.Sleep(2 * time.Second)
		//		}
	}
	{
		no := 3
		var wg sync.WaitGroup
		for i := 0; i < no; i++ {
			wg.Add(1)
			go process(i, &wg)
		}
		wg.Wait()
		fmt.Println("all go routines finished")
	}
}
