package main

import (
	"log"
	"sync"
)

func main() {
	maxWorkers := 5
	resultCh := make(chan int)

	var wg sync.WaitGroup

	go AggregatorGopher(resultCh)

	for v := 1; v <= maxWorkers; v++ {
		wg.Add(1)
		go SquarerGopher(v, &wg, resultCh)
	}

	wg.Wait()
	close(resultCh)
}

func SquarerGopher(v int, wg *sync.WaitGroup, resultCh chan<- int) {
	defer wg.Done()
	log.Printf("Squaring: %d", v)
	squared := v * v
	resultCh <- squared
}

func AggregatorGopher(resultCh <-chan int) {
	total := 0
	for val := range resultCh {
		log.Printf("Aggregator received: %d", val)
		total += val
		log.Printf("Current total: %v", total)
	}
	log.Printf("Aggregator finished. Total sum: %d", total)
}
