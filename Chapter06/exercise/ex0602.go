package main

import "fmt"

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println("Inside UpdateSlice:", ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println("Inside GrowSlice:", ss)
}

func main() {
	s := []string{"A", "B", "C"}
	fmt.Println("Before UpdateSlice:", s)
	UpdateSlice(s, "Z")
	fmt.Println("After UpdateSlice:", s)

	fmt.Println("Before GrowSlice:", s)
	GrowSlice(s, "Y")
	fmt.Println("After GrowSlice:", s)
}
