package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main() {

	mainChannel := make(chan int)

	n := 100
	go Generate(mainChannel, n)

	primeChannel := make(chan int)

	var wg sync.WaitGroup

	primes := make([]int, 0)

	go func() {
		for i := range primeChannel {
			primes = append(primes, i)
		}
	}()

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go filter(mainChannel, primeChannel, &wg)
	}

	wg.Wait()

	sort.Ints(primes)

	for i := range primes {
		fmt.Println(i)
	}

	max := 10
	pp := PerfectPowers(primes, max)

	for i := range pp {
		fmt.Println(i)
	}

	fmt.Println("dank")
}

func PerfectPowers(primes []int, maxPower int) []string {
	perfectPowers := make([]string, 0)

	for i := range primes {
		for index := range primes {
			prime := primes[index]
			sum := 0

			for i := 0; i < len(primes); i++ {
				sum += primes[i]
			}

			for power := 2; power < maxPower+1; power++ {
				if isPower(sum, power) {
					s := string(prime) + ":" + string(primes[i]) + " = " + string(sum) + " = " + string(getNthroot(sum, power)) + "**" + string(power)
					perfectPowers = append(perfectPowers, s)
				}
			}
		}
	}

	return perfectPowers
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

func isPower(i, base int) bool {
	diff := math.Mod(math.Pow(float64(i), (1/float64(base))), 1)

	return diff == 0.0 || diff > 0.9999999
}

func getNthroot(a, b int) int {
	s := math.Pow(float64(a), 1/float64(b))
	return int(s)
}
