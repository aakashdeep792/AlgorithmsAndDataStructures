package code

import (
	"fmt"
	"sync"
)

/*
Go program that calculates the sum of all even numbers from 1 to 100 using one hundred goroutines and channels:

*/

func CalculateSum(n int) string {
	var wg sync.WaitGroup
	ch := make(chan int)
	inc := n / 4

	for i := 1; i <= n; {
		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()

			fmt.Println(s, e)
			sum := 0
			for i := s; i <= e; i++ {
				if i%2 == 0 {
					sum += i
				}
			}

			ch <- sum

		}(i, min(i+inc, n))

		i = i + inc + 1
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for s := range ch {
		sum += s
	}

	fmt.Println(sum)

	return "success"
}
