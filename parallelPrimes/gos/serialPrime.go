package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/sgrade/parallelPrimes/gos/helpers"
)

func main() {

	args := os.Args
	lower, _ := strconv.ParseInt(args[1], 10, 64)
	upper, _ := strconv.ParseInt(args[2], 10, 64)
	maxPower, _ := strconv.ParseInt(args[3], 10, 64)
	fileArg := args[4]

	primes := make([]int, 0)

	start := time.Now()

	for i := int(lower); i <= int(upper); i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}

	file := helpers.FileCreate(fileArg)
	defer file.Close()

	PerfectPowers(primes, int(maxPower), file)

	fmt.Println(time.Since(start))

}

func IsPrime(n int) bool {
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

			for power := 2; power <= maxPower; power++ {
				if helpers.IsPower(sum, power) {
					s := strconv.FormatInt(int64(prime), 10) + ":" + strconv.FormatInt(int64(primes[elem]), 10) + " = " + strconv.FormatInt(int64(sum), 10) + " = " + strconv.FormatInt(int64(helpers.GetNthroot(sum, power)), 10) + "**" + strconv.FormatInt(int64(power), 10)
					io.WriteString(fp, s+"\n")
				}
			}
		}
	}
}
