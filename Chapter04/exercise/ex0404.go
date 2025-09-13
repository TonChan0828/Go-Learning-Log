package main

import "fmt"

func main() {
	total := 0
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Println(total)
}
