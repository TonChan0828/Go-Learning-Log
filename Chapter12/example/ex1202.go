package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		v := 1
		fmt.Printf("木の下でch1へ%vを入れる\n", v)
		ch1 <- v
	}()

	go func() {
		v := 1
		fmt.Printf("木の下でch2へ%vを入れる\n", v)
		ch2 <- v
	}()

	go func() {
		v := 1
		fmt.Printf("木の下でch3へ%vを入れる\n", v)
		ch3 <- v
	}()

	go func() {
		v := 1
		fmt.Printf("木の下でch4へ%vを入れる\n", v)
		ch4 <- v
	}()

	select {
	case v := <-ch1:
		fmt.Println("ch1から読み込み:", v)
	case v := <-ch2:
		fmt.Println("ch2から読み込み:", v)
	case v := <-ch3:
		fmt.Println("ch3から読み込み:", v)
	case v := <-ch4:
		fmt.Println("ch4から読み込み:", v)
	}
}
