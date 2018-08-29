package main

import (
	"fmt"
)

func main() {
	n := 2000

	count := 0

	for i := 0; ; i++ {
		if isPrime(i) {
			fmt.Println(i)
			count++
		}

		if count == n {
			break
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
