package main

import (
	"fmt"
	"sync"
	"strconv"
	"math"
	"os"
	"log"
	"io"
)

var waitGroup sync.WaitGroup

func main() {
	fmt.Println("Starting the application...")
	data := make(chan int)

	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal("Cannot create file.", err)
	}

	defer file.Close()

	array := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker(data, 4, file)
	}

	for i := 0; i < len(array); i++ {
		data <- array[i]
	}

	close(data)

	waitGroup.Wait()
}

func worker(data chan int, maxpower int, fp *os.File) {
	defer func() {
		waitGroup.Done()
	}()

	primebuf := make([]int, 0)

	for {
		value, ok := <-data

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

			for power := 2; power <= maxpower; power++ {
				if isPower(sum, power) {
					s := strconv.FormatInt(int64(prime), 10) + ":" + strconv.FormatInt(int64(value), 10) + " = " + strconv.FormatInt(int64(sum), 10) + " = " + strconv.FormatInt(int64(getNthroot(sum, power)), 10) + "**" + strconv.FormatInt(int64(power), 10)
					io.WriteString(fp, s+"\n")
				}
			}
		}
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
