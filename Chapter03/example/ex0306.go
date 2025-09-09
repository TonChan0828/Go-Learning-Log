package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	y = append(y, "z")
	y = append(y, "1")
	y = append(y, "2")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println(cap(x), cap(y))
}
