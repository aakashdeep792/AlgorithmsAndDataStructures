package code

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

/*

Write a Go program that concurrently counts the occurrences of each word in a given list of text files.
Implement a function WordCount(filename string) map[string]int

*/

var (
	globalFreq = make(map[string]int)
	re         = regexp.MustCompile(`[\t\n ]`)
)

func wordCounter(wg *sync.WaitGroup, m *sync.Mutex, file string) {
	defer wg.Done()

	bytes, err := os.ReadFile(file)
	if err != nil {
		return
	}
	str := string(bytes)
	// fmt.Println(str)

	freq := make(map[string]int)

	// str = strings.ReplaceAll(str, "[\t\n ]", " ")
	str = re.ReplaceAllString(str, " ")
	words := strings.Split(str, " ")

	for _, w := range words {
		if w == " " || w == "" {
			continue
		}
		freq[w]++
	}

	m.Lock()
	defer m.Unlock()
	for k, v := range freq {
		globalFreq[k] += v
	}
	// fd,err:= os.OpenFile(file,os.O_RDONLY,0644)
	// if err!= nil{
	// 	return
	// }

	// while fd.
}

func WordCountInFile(files ...string) string {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, f := range files {
		wg.Add(1)
		go wordCounter(&wg, &m, f)
	}
	wg.Wait()

	for k, v := range globalFreq {
		fmt.Println(k, " = ", v)
	}

	return "success"
}
