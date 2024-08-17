package main

import (
	"fmt"
)

type IntOrFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func main() {
	fmt.Println(doubler(3.14))
	fmt.Println(doubler(9000))
}

func doubler[T IntOrFloat](t T) T {
	return t * 2
}
