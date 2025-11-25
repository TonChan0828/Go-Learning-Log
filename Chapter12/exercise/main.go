package main

import "fmt"

func ProcessData() {
	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i*100 + 1
		}
		close(ch2)
	}()

	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				count--
				break
			}
			fmt.Println(v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Println(v)
		}
	}
}

func main() {
	ProcessData()
}
