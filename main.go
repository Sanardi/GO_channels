package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		time.Sleep(5 * time.Millisecond)
		//fmt.Println(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
