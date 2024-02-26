package producer

import (
	"fmt"
	"sync"
)

type Producer interface {
	Produce(val int)
}

type ItemProducer struct {
	ch chan string //buffer channel
	wg *sync.WaitGroup
}

func NewItemProducer(ch chan string, wg *sync.WaitGroup) Producer {
	return &ItemProducer{ch: ch, wg: wg}
}

// Produce a Number from 0 to val
func (i *ItemProducer) Produce(val int) {
	for count := 0; count <= val; count++ {
		i.ch <- fmt.Sprint(count)
	}

	close(i.ch)
}
