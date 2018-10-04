package main

import (
	"fmt"
	"sync"
)

func main() {
	mainChannel := make(chan int)

	n := 530000
	go Generate(mainChannel, n)

	primeChannel := make(chan int)

	var wg sync.WaitGroup

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go filter(mainChannel, primeChannel, &wg)
	}

	for {
		i, more := <-primeChannel
		if more {
			fmt.Println(i)
		} else {
			fmt.Println()
		}
	}

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func filter(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := <-in
	if isPrime(i) {
		out <- i
	}
}

func Generate(ch chan<- int, thresh int) {
	for i := 2; i <= thresh; i++ {
		ch <- i
	}
}
