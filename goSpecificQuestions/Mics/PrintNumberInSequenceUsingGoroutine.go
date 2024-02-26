package main

import (
	"fmt"
	"sync"
)

func PrintAndIncrNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-ch
	fmt.Println(val)
	ch <- val + 1
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	val := 50
	ch <- 1
	wg.Add(val)

	// spawn multiple go routine
	for count := 0; count < val; count++ {
		go PrintAndIncrNum(ch, wg)
	}

	wg.Wait()
}
