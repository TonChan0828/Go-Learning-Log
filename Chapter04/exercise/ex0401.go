package main

import (
	"fmt"
	"math/rand"
)

func main() {
	v := []int{}
	for i := 0; i < 100; i++ {
		v = append(v, rand.Intn(100))
	}
	fmt.Println(v)
}
