package consumer

import (
	"fmt"
	"sync"
)

type Consumer interface {
	Consume()
}

type ItemConsumer struct {
	ch chan string
	wg *sync.WaitGroup
}

func NewItemConsumer(ch chan string, wg *sync.WaitGroup) Consumer {
	return &ItemConsumer{ch: ch, wg: wg}
}

// Consume() print all the numbers consumed
func (i *ItemConsumer) Consume() {
	defer i.wg.Done()

	var output []string

	for val := range i.ch {
		output = append(output, val)
	}

	fmt.Println("Consumer Output", output)
}
