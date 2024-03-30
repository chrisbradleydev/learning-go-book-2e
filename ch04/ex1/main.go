package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var slice []int
	for i := 0; i < 100; i++ {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println(len(slice))
	fmt.Println(slice)
}
