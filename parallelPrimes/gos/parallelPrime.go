package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {

	mainChannel := make(chan int)

	n := 100

	go Generate(mainChannel, n)

	primeChannel := make(chan int)

	var wg sync.WaitGroup

	primes := make([]int, 0)

	start := time.Now()
	go func() {
		for i := range primeChannel {
			//primes = append(primes, i)
			fmt.Println(i)
		}
	}()

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go filter(mainChannel, primeChannel, &wg)
	}

	wg.Wait()
	close(primeChannel)

	fmt.Println(time.Since(start))
	sort.Ints(primes)

	max := 10

}

func PerfectPowers(primes []int, maxPower int) {
	perfectPowers := make([]string, 0)
	primebuf := make([]int, 0)

	for elem := range primes {
		primebuf = append(primebuf, primes[elem])

		for index := 0; index < len(primebuf); index++ {
			prime := primebuf[index]
			sum := 0

			for i := index; i < len(primebuf); i++ {
				sum += primebuf[i]
			}

			for power := 2; power < maxPower+1; power++ {
				if isPower(sum, power) {
					s := strconv.FormatInt(int64(prime), 10) + ":" + strconv.FormatInt(int64(primes[elem]), 10) + " = " + strconv.FormatInt(int64(sum), 10) + " = " + strconv.FormatInt(int64(getNthroot(sum, power)), 10) + "**" + strconv.FormatInt(int64(power), 10)
					perfectPowers = append(perfectPowers, s)
				}
			}
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

func isPower(num, power int) bool {
	err := math.Pow(float64(num), 1/float64(power))

	t := math.Abs(err - math.Round(err))

	if t < 0.0000001 {
		return true
	}
	return false
}

func getNthroot(a, b int) int {
	s := math.Pow(float64(a), 1/float64(b))

	return int(math.Round(s))
}
