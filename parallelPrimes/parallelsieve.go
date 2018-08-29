// A concurrent prime sieve

package main

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 2000; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	fmt.Println()
}
