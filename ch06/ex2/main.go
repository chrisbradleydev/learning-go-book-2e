package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println("in main, before UpdateSlice:", s)
	UpdateSlice(s, "d")
	fmt.Println("in main, after UpdateSlice:", s)
	GrowSlice(s, "e")
	fmt.Println("in main, after GrowSlice:", s)
}

func UpdateSlice(s []string, val string) {
	s[len(s)-1] = val
	fmt.Println("in UpdateSlice:", s)
}

func GrowSlice(s []string, val string) {
	s = append(s, val)
	fmt.Println("in GrowSlice:", s)
}
