package main

import (
	"fmt"
)

func main() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(len(runes))
	fmt.Println(string(runes[3]))
	fmt.Println(string(runes[9]))
}
