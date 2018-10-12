package main

import (
	"fmt"
	"sync"
)

func main() {
	mainChannel := make(chan int)

	n := 100000
	go Generate(mainChannel, n)

	primeChannel := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for i := range primeChannel {
			fmt.Println(i)
		}
	}()

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go filter(mainChannel, primeChannel, &wg)
	}

	wg.Wait()
	close(primeChannel)
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
