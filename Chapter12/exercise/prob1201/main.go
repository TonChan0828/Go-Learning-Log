package main

import "sync"

func ProcessData() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i*100 + i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		for v := range ch {
			println(v)
		}
	}()

	wg.Wait()
}

func main() {
	ProcessData()
}
