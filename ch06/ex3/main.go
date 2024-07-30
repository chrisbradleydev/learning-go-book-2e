package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	entries := 10_000_000
	people := make([]Person, 0, entries)
	for i := 0; i < entries; i++ {
		people = append(people, MakePerson("Bob", "Barker", 80))
	}
	fmt.Println(len(people))
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
