package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

func main() {

	mainChannel := make(chan int)

	n := 10
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
		fmt.Println(primes[i])
	}

	max := 4

	pp := PerfectPowers(primes, max)

	for i := range pp {
		fmt.Println(string(pp[i]))
	}
}

func PerfectPowers(primes []int, maxPower int) []string {
	perfectPowers := make([]string, 0)
	primebuf := make([]int, 0)

	for elem := range primes {
		primebuf = append(primebuf, primes[elem])

		for index := range primebuf {
			prime := primebuf[index]
			sum := 0

			for i := index; i < len(primebuf); i++ {
				sum += primebuf[i]
			}

			fmt.Println("Sum =", sum)

			for power := 2; power < maxPower+1; power++ {
				if isPower(sum, power) {
					s := strconv.FormatInt(int64(prime), 10) + ":" + strconv.FormatInt(int64(primes[elem]), 10) + " = " + strconv.FormatInt(int64(sum), 10) + " = " + strconv.FormatInt(int64(power), 10) + "**" + strconv.FormatInt(int64(getNthroot(sum, power)), 10)
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

func isPower(num, base int) bool {
	k := math.Log(float64(num)) / math.Log(float64(base))

	t := math.Abs(k - math.Round(k))

	if t < 0.0000001 {
		return true
	}
	return false
}

func getNthroot(a, b int) int {
	s := math.Pow(float64(a), 1/float64(b))

	return int(math.Round(s))
}
