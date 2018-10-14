package helpers

import (
	"math"
	"sync"
	"os"
	"log"
)

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

func Filter(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := <-in
	if IsPrime(i) {
		out <- i
	}
}

func Generate(ch chan<- int, thresh int) {
	for i := 2; i <= thresh; i++ {
		ch <- i
	}
}

func IsPower(num, power int) bool {
	err := math.Pow(float64(num), 1/float64(power))

	t := math.Abs(err - math.Round(err))

	if t < 0.0000001 {
		return true
	}
	return false
}

func GetNthroot(a, b int) int {
	s := math.Pow(float64(a), 1/float64(b))

	return int(math.Round(s))
}

func FileCreate(name string) *os.File {
	file, err := os.Create(name)

	if err != nil {
		log.Fatal("Cannot create file.", err)
	}

	return file
}