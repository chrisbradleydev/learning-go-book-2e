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

	for _, v := range slice {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind.")
		}
	}
}
