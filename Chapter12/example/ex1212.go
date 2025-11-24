package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		doThing1()
	}()

	go func() {
		defer wg.Done()
		doThing2()
	}()

	go func() {
		defer wg.Done()
		doThing3()
	}()

	wg.Wait()
}

func doThing1() {
	time.Sleep(getSec(1))
	fmt.Println("doThing1 Done")
}

func doThing2() {
	time.Sleep(getSec(1))
	fmt.Println("doThing2 Done")
}

func doThing3() {
	time.Sleep(getSec(1))
	fmt.Println("doThing3 Done")
}

func getSec(id int) time.Duration {
	sec := rand.Intn(6 + 1)
	fmt.Printf("doThing%d: %v sec\n", id, sec)
	return time.Duration(time.Duration(sec) * time.Second)
}
