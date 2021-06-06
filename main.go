package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		msg, ok := <-ch
		fmt.Println(msg, ok)
		//fmt.Println(<-ch)
		//close(ch)
		//fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		//ch <- 42
		//time.Sleep(5 * time.Millisecond)
		//fmt.Println(ch)
		close(ch)
		//ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
