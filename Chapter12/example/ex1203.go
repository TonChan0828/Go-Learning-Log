package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "ゴルーチンから送信した文字列"
		v1 := <-ch2
		fmt.Println(v1)
	}()

	ch2 <- "mainから送信した文字列"
	v2 := <-ch1
	fmt.Println(v2)
}
