package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"Marcus", "Aurelius", 1}
	emp2 := Employee{
		firstName: "Nikola",
		lastName: "Tesla",
		id: 2,
	}
	var emp3 Employee
	emp3.firstName = "Akira"
	emp3.lastName = "Toriyama"
	emp3.id = 3

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
