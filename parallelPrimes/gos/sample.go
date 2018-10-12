package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(rev(124, 3))
}

func isPower(num, base int) bool {
	k := math.Log(float64(num)) / math.Log(float64(base))

	t := math.Abs(k - math.Round(k))

	if t < 0.0000001 {
		return true
	}
	return false
}

func rev(num, power int) bool {
	err := math.Pow(float64(num), 1/float64(power))

	t := math.Abs(err - math.Round(err))

	if t < 0.0000001 {
		return true
	}
	return false
}
