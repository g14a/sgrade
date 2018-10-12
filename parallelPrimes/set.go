package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(k int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- k
		}(i, &wg)
	}

	wg.Wait()
	close(ch)
}
