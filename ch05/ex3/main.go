package main

import (
	"fmt"
)

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Cloud"))
	fmt.Println(helloPrefix("Squall"))
	fmt.Println(helloPrefix("Zidane"))
	fmt.Println(helloPrefix("Tidus"))
}

func prefixer(prefix string) func(string) string {
	return func(subject string) string {
		return prefix + " " + subject
	}
}
