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

	var wg, pwg sync.WaitGroup
	primes := make([]int, 0)

	args := os.Args
	lower, _ := strconv.ParseInt(args[1], 10, 64)
	upper, _ := strconv.ParseInt(args[2], 10, 64)
	maxPower, _ := strconv.ParseInt(args[3], 10, 64)
	fileArg := args[4]

	file := helpers.FileCreate(fileArg)
	defer file.Close()

	start := time.Now()

	go helpers.Generate(int(lower), mainChannel, int(upper))

	go func() {
		for i := range primeChannel {
			primes = append(primes, i)
		}
	}()

	for i := int(lower); i <= int(upper); i++ {
		wg.Add(1)
		go helpers.Filter(mainChannel, primeChannel, &wg)
	}

	end := time.Since(start)

	wg.Wait()
	close(primeChannel)

	sort.Ints(primes)

	again := time.Now()

	PerfectPowers(primes, int(maxPower), file)

	done := time.Since(again)
	pwg.Wait()

	fmt.Println(end + done)
}

func PerfectPowers(primes []int, maxPower int, fp *os.File) {
	primebuf := make([]int, 0)

	for elem := range primes {
		primebuf = append(primebuf, primes[elem])
		dep := primes[elem]

		for index := 0; index < len(primebuf); index++ {
			prime := primebuf[index]
			sum := 0
			t := dep

			for i := index; i < len(primebuf); i++ {
				sum += primebuf[i]
			}

			var wg sync.WaitGroup

			go func(wg *sync.WaitGroup) {
				wg.Add(1)
				defer wg.Done()
				for power := 2; power <= maxPower; power++ {
					if helpers.IsPower(sum, power) {
						fmt.Println(sum)
						s := strconv.Itoa(prime) + ":" + strconv.Itoa(t) + " = " + strconv.Itoa(sum) + " = " + strconv.Itoa(helpers.GetNthroot(sum, power)) + "**" + strconv.Itoa(power)
						io.WriteString(fp, s+"\n")
					}
				}
			}(&wg)
		}
	}
}
