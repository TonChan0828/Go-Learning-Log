package main

import "fmt"

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func main() {
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result1 := myFuncVariable("Hello")
	fmt.Println("result1=", result1)

	myFuncVariable = f2
	result2 := myFuncVariable("Hello")
	fmt.Println("result2=", result2)
}
