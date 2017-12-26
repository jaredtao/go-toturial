package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}
func incrementWithChan(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	{
		//		var wg sync.WaitGroup
		//		var mutex sync.Mutex
		//		for i := 0; i < 1000; i++ {
		//			wg.Add(1)
		//			go increment(&wg, &mutex)
		//		}
		//		wg.Wait()
	}
	{
		var wg sync.WaitGroup
		ch := make(chan bool, 1)
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go incrementWithChan(&wg, ch)
		}
		wg.Wait()
	}
	fmt.Println("final value of x", x)
}
