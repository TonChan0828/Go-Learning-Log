package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "ゴルーチンから送信した文字列"
		fromMain := <-ch2
		fmt.Println("無名関数:", fromMain)
	}()

	var fromGoroutine string
	select {
	case ch2 <- "mainから送信した文字列":
	case fromGoroutine = <-ch1:
	}
	fmt.Println("main:", fromGoroutine)
}
