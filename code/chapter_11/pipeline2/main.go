package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for num := 0; num <= 10; num++ {
		out <- num
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}
