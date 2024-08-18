package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	total := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("total:", total, "number of iterations:", count, ctx.Err())
			return
		default:
		}
		newNum := rand.Intn(100_000_000)
		if newNum == 1_234 {
			fmt.Println("total:", total, "number of iterations:", count, "got 1,234")
			return
		}
		total += newNum
		count++
	}
}
