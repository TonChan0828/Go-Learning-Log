package main

import "fmt"

func prefixer(s string) func(string) string {
	return func(s2 string) string {
		return s + " " + s2
	}
}

func main() {
	helloPrefix := prefixer(("Hello"))
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
