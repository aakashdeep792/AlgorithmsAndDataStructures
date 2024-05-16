package code

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type csvReader struct {
	reader *csv.Reader
	ch     chan info
}

type info struct {
	line    int
	details []string
}

type iStorage interface {
	Set(line int, details []string)
}

func NewRecordStore() iStorage {
	return &RecordStore{
		dict: make(map[int][]string),
	}
}

type RecordStore struct {
	dict map[int][]string
	mu   sync.RWMutex
}

func (s *RecordStore) Set(line int, record []string) {
	// tmp := "line_" + strconv.Itoa(line)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dict[line] = record
}

// var (
// 	lineDict = make(map[int][]string)
// 	// recordStore = make([]string, 0)
// )

func NewCSVReader(reader *csv.Reader, ch chan info) *csvReader {
	return &csvReader{
		reader: reader,
		ch:     ch,
	}
}

func ReadCSV(file string) string {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	reader := csv.NewReader(fd)

	ch := make(chan info)
	r := NewCSVReader(reader, ch)
	var wg sync.WaitGroup
	rs := NewRecordStore()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go r.Read(&wg)
	}

	// close channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// update the line number in dict
	// in this case mutex is not require
	for i := range ch {
		rs.Set(i.line, i.details)
	}

	printStore(rs)

	return "success"
}

func printStore(s iStorage) {
	rs := s.(*RecordStore)
	count := 0
	for k, v := range rs.dict {
		if count > 5 {
			return
		}
		fmt.Println(k, v)
		count++
	}
}

func (c *csvReader) Read(wg *sync.WaitGroup) {
	wg.Done()

	defer func(){
		if r:= recover(); r!= nil{
			fmt.Println("worker Paniced:",r)
		}
	}()

	for {

		record, err := c.reader.Read()
		if err != nil {
			fmt.Println("Worker stopped:", err)
			return
		}

		// process string
		// get line number
		lineNo := c.reader.InputOffset()

		c.ch <- info{line: int(lineNo), details: record}
	}

}
