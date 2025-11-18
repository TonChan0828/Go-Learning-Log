package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer func() {
			close(ch)
			fmt.Println("chをクローズした")

		}()

		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("mainを終了")
}
