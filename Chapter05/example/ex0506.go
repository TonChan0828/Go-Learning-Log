package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(num int, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return num, denom, errors.New("denom can not be zero")
	}
	result = num / denom
	remainder = num % denom
	return result, remainder, err

}

func callDivAndRemainder(num int, denom int) {
	x, y, z := divAndRemainder(num, denom)
	if z != nil {
		fmt.Print(x, "+", y, ":")
		fmt.Println(z)
		os.Exit(1)
	}
	fmt.Print(num, "+", denom, "=", x, "…", y, "\n")
}

func main() {
	callDivAndRemainder(5, 2)
	callDivAndRemainder(10, 0)
}
