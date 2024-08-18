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
		message := fmt.Sprintf("record %d: %+v", count, emp)
		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				allErrors := err.Unwrap()
				var messages []string
				for _, e := range allErrors {
					messages = append(messages, processError(e, emp))
				}
				message = message + " allErrors: " + strings.Join(messages, ", ")
			default:
				message = message + " error: " + processError(err, emp)
			}
		}
		fmt.Println(message)
	}
}

func ValidateEmployee(e Employee) error {
	var allErrors []error
	if len(e.ID) == 0 {
		allErrors = append(allErrors, EmptyFieldError{FieldName: "ID"})
	} else if !validID.MatchString(e.ID) {
		allErrors = append(allErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		allErrors = append(allErrors, EmptyFieldError{FieldName: "FirstName"})
	}
	if len(e.LastName) == 0 {
		allErrors = append(allErrors, EmptyFieldError{FieldName: "LastName"})
	}
	if len(e.Title) == 0 {
		allErrors = append(allErrors, EmptyFieldError{FieldName: "Title"})
	}
	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}

func processError(err error, emp Employee) string {
	var fieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("invalid ID: %s", emp.ID)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}
