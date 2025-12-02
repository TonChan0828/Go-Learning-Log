package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	total := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("total:", total, " iterations:", count, " Timeout")
			return
		default:
		}
		newNum := rand.Intn(100_000_000)
		if newNum == 1_234 {
			fmt.Println("total:", total, " iterations:", count, " Success")
			return
		}
		total += newNum
		count++
	}
}
