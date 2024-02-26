package main

import (
	"algoDS/goSpecificQuestions/producerConsumer/consumer"
	"algoDS/goSpecificQuestions/producerConsumer/producer"
	"fmt"
	"sync"
)

func main() {
	var (
		wg  sync.WaitGroup
		val int
	)

	ch := make(chan string, 5)

	// scan from console
	fmt.Println("Enter a number eg 23")
	fmt.Scanln(&val)

	c := consumer.NewItemConsumer(ch, &wg)
	p := producer.NewItemProducer(ch, &wg)

	wg.Add(1) // if this statemnet is added inside the consumer then wg.Wait() will be execute before the wg.Add() (in some scenario) the program will run without any output.
	go c.Consume()
	go p.Produce(val)

	wg.Wait()
}
