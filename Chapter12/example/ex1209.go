package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processChannel(ch chan int) []int {
	const maxConc = 10
	results := make(chan int, maxConc)
	for i := 0; i <= maxConc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	fmt.Println("ゴルーチン　起動完了")

	var out []int
	for i := 0; i < maxConc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	returnVal := v * v
	sleepSec := rand.Intn(3)

	fmt.Println("process:", v, returnVal, sleepSec)
	time.Sleep(time.Duration(sleepSec) * time.Second)
	return returnVal
}

func main() {
	ch := make(chan int)

	var result []int
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	result = processChannel(ch)
	fmt.Printf("result: %d\n", result)
}
