package models

import (
	"errors"
	"fmt"
)

type Person struct {
	Title     string
	FirstName string
	LastName  string
}

func (p *Person) Saludar() error {
	if p.FirstName == "" {
		return errors.New("No name")
	}

	fmt.Printf("Saludos %s %s", p.Title, p.FirstName)
	return nil
}
