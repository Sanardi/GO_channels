package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//if msg, ok := <-ch; ok {
		//fmt.Println(msg, ok)
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
		//fmt.Println(<-ch)
		//close(ch)
		//fmt.Println(<-ch)
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		//ch <- 0
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//time.Sleep(5 * time.Millisecond)
		//fmt.Println(ch)
		close(ch)
		//ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
