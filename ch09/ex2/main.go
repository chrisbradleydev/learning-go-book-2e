package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type EmptyFieldError struct {
	FieldName string
}

func (fe EmptyFieldError) Error() string {
	return fe.FieldName
}

//go:embed employees.txt
var data string

var (
	ErrInvalidID = errors.New("invalid ID")
	validID      = regexp.MustCompile(`\w{4}-\d{3}`)
)

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		var fieldErr EmptyFieldError
		if err != nil {
			if errors.Is(err, ErrInvalidID) {
				fmt.Printf("record %d: %+v error: invalid ID: %s\n", count, emp, emp.ID)
			} else if errors.As(err, &fieldErr) {
				fmt.Printf("record %d: %+v error: empty field %s\n", count, emp, fieldErr.FieldName)
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			}
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return EmptyFieldError{FieldName: "ID"}
	}
	if !validID.MatchString(e.ID) {
		return ErrInvalidID
	}
	if len(e.FirstName) == 0 {
		return EmptyFieldError{FieldName: "FirstName"}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldError{FieldName: "LastName"}
	}
	if len(e.Title) == 0 {
		return EmptyFieldError{FieldName: "Title"}
	}
	return nil
}
