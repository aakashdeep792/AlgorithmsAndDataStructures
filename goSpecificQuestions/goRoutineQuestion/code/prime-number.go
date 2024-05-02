package code

import (
	"fmt"
	"math"
	"sync"
)

func isPrime(wg *sync.WaitGroup, ch chan int, n int) {
	defer wg.Done()
	tmp := int(math.Sqrt(float64(n)))

	for i := 2; i <= tmp; i++ {
		if n%i == 0 {
			return
		}
	}

	ch <- n
}

func CheckPrime(n int) string {
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := 2; i <= n; i++ {
		wg.Add(1)
		go isPrime(&wg, ch, i)
	}
	// wait and close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	return "success"
}
