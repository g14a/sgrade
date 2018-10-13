package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1000000

	start := time.Now()

	for i := 0; i < n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	fmt.Println(time.Since(start))
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
