package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sort"
)

func main() {

	mainChannel := make(chan int)

	n := 40

	go Generate(mainChannel, n)

	primeChannel := make(chan int)

	var wg,pwg sync.WaitGroup
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
	close(primeChannel)

	sort.Ints(primes)
	
	max := 5
	
	sortedPrime := make(chan int)

	for i:=0;i<len(primes);i++ {
		pwg.Add(1)
		go PerfectPowers(sortedPrime, max, &pwg)
	}

	for i:=0;i<len(primes);i++ {
		sortedPrime <- primes[i]
	}

	close(sortedPrime)
	pwg.Wait()

}

func PerfectPowers(primeChannel chan int, maxPower int, wg *sync.WaitGroup) {

	defer wg.Done()
	primebuf := make([]int, 0)

	for {
		value, ok := <-primeChannel
		if !ok {
			break
		}

		primebuf = append(primebuf, value)

		for index := 0; index < len(primebuf); index++ {
			prime := primebuf[index]
			sum := 0

			for i := index; i < len(primebuf); i++ {
				sum += primebuf[i]
			}

			for power := 2; power <= maxPower; power++ {
				if isPower(sum, power) {
					s := strconv.FormatInt(int64(prime), 10) + ":" + strconv.FormatInt(int64(value), 10) + " = " + strconv.FormatInt(int64(sum), 10) + " = " + strconv.FormatInt(int64(getNthroot(sum, power)), 10) + "**" + strconv.FormatInt(int64(power), 10)
					fmt.Println(s)
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
