package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/sgrade/parallelPrimes/gos/helpers"
)

func main() {

	mainChannel := make(chan int)
	primeChannel := make(chan int)
	//sortedPrime := make(chan int)

	var wg, pwg sync.WaitGroup
	primes := make([]int, 0)

	n := 50000

	file := helpers.FileCreate("file.txt")
	defer file.Close()

	start := time.Now()

	go helpers.Generate(mainChannel, n)

	go func() {
		for i := range primeChannel {
			primes = append(primes, i)
			//fmt.Println(i)
		}
	}()

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go helpers.Filter(mainChannel, primeChannel, &wg)
	}

	end := time.Since(start)

	wg.Wait()
	close(primeChannel)

	sort.Ints(primes)

	again := time.Now()

	maxPower := 5
	PerfectPowers(primes, maxPower, file)

	pwg.Wait()

	fmt.Println(end + time.Since(again))
}

func PerfectPowers(primes []int, maxPower int, fp *os.File) {
	primebuf := make([]int, 0)

	for elem := range primes {
		primebuf = append(primebuf, primes[elem])

		for index := 0; index < len(primebuf); index++ {
			prime := primebuf[index]
			sum := 0

			for i := index; i < len(primebuf); i++ {
				sum += primebuf[i]
			}

			var wg sync.WaitGroup

			go func(wg *sync.WaitGroup) {
				wg.Add(1)
				defer wg.Done()
				for power := 2; power <= maxPower; power++ {
					if helpers.IsPower(sum, power) {
						s := strconv.FormatInt(int64(prime), 10) + ":" + strconv.FormatInt(int64(primes[elem]), 10) + " = " + strconv.FormatInt(int64(sum), 10) + " = " + strconv.FormatInt(int64(helpers.GetNthroot(sum, power)), 10) + "**" + strconv.FormatInt(int64(power), 10)
						io.WriteString(fp, s+"\n")
					}
				}
			}(&wg)
		}
	}
}
