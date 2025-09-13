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

	for _, v := range v {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
		}else if v%2==0{
			fmt.Println("Two!")
		}else if v%3==0{
			fmt.Println("Three!")
		}else{
			fmt.Println("Never mind")
		}
	}
}
