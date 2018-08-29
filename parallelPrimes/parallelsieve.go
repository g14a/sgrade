package main

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	input := make(chan int)
	go Generate(input)
	for i := 0; i < 2000; i++ {
		prime := <-input
		fmt.Println(prime)
		output := make(chan int)
		go Filter(input, output, prime)
		input = output
	}

	fmt.Println()
}
