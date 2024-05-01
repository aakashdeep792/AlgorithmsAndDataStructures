package code

import (
	"fmt"
	"sync"
)

func print(wg *sync.WaitGroup, ch chan int) {
	tmp := <-ch
	tmp++
	fmt.Println(tmp)
	ch <- tmp
	wg.Done()
}

// method 1 not efficient
func PrintNumber1(n int) string {
	var wg sync.WaitGroup
	fmt.Println("code starting")
	ch := make(chan int)

	fmt.Println("value inserted")

	for i := 0; i < n; i++ {
		wg.Add(1)
		go print(&wg, ch)
	}

	ch <- 0
	wg.Add(-1) // the goroutine that is running last will have value in channel, so last go routine will be blocked
	wg.Wait()

	return "success"
}

// output will be unsorted
func PrintNumber2(n int) string {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	return "success"
}

func PrintNumber3(n int) string {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Start 100 goroutines to generate numbers concurrently
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n // Send the number to the channel
		}(i + 1)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Read numbers from the channel and print them in sorted order
	sortedNumbers := make([]int, 100)
	for num := range ch {
		sortedNumbers[num-1] = num
	}

	// Print the sorted numbers
	for _, num := range sortedNumbers {
		fmt.Println(num)
	}

	return "success"
}

// using mutex
func PrintNumber4(n int) string {
	var nums int
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			nums++
			fmt.Println(nums)
			wg.Done()

		}()
	}

	wg.Wait()

	return "success"
}
