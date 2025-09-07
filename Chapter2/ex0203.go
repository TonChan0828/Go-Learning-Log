package main

import (
	"fmt"
)

func main() {
	var b byte = 255
	b += 1
	fmt.Println(b)

	var smallI int32 = 1<<31 - 1
	smallI += 1
	fmt.Println(smallI)

	var bigI uint64 = 1<<64 - 1
	bigI += 1
	fmt.Println(bigI)
}
