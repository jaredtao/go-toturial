package main

import "fmt"

func hello(c chan bool) {
	fmt.Println("Hello chan")
	c <- true
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	digit := 0
	for number != 0 {
		digit = number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}
func calcCube(number int, cubeop chan int) {
	sum := 0
	digit := 0
	for number != 0 {
		digit = number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func digits(number int, dchan chan int) {
	digit := 0
	for number != 0 {
		digit = number % 10
		dchan <- digit
		number /= 10
	}
	close(dchan)
}
func calcSquares_2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}
func main() {
	c := make(chan bool)
	go hello(c)
	fmt.Println("main", <-c)

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCube(number, cubech)
	sqrs := <-sqrch
	cubes := <-cubech
	//	sqrs, cubes := <-sqrch, <-cubech
	fmt.Println("out", sqrs, cubes)
	{
		number := 589
		sqrch := make(chan int)
		go calcSquares_2(number, sqrch)
		sqrs := <-sqrch
		fmt.Println("calcSquares_2", sqrs)
	}

	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}
	for v := range ch {
		fmt.Println("Received", v)
	}

}
