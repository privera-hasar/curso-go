package main

import (
	"curso/abstractions"
	"curso/models"
	"fmt"
)

func main() {

	var p abstractions.IPerson

	p = &models.Person{
		Title:     "Don",
		FirstName: "Juan",
		LastName:  "Perez",
	}

	err := p.Saludar()
	if err != nil {
		fmt.Print(err.Error())
	}
}

func sample1() {
	fmt.Println("Hello world")
}

func sample2() {
	var message string
	message = "Hello world"
	fmt.Println(message)
}

func sample3() {
	message := "Hello world"
	fmt.Println(message)
}

func sample5() {
	shouldPrint := true
	times := 3

	if shouldPrint {
		for i := 1; i <= times; i++ {
			message := "Hello world"
			fmt.Println(message)
		}
	}
}

func sample6() {
	shouldPrint := true
	times := 3

	if shouldPrint {
		i := 1
		for i <= times {
			message := "Hello world"
			fmt.Println(message)
			i++
		}
	}
}

func sample7() {

	p := &models.Person{
		Title:     "Don",
		FirstName: "Juan",
		LastName:  "Perez",
	}

	fmt.Printf("Saludos %s %s", p.Title, p.FirstName)
}

func sample8() {

	p := &models.Person{
		Title:     "Don",
		FirstName: "Juan",
		LastName:  "Perez",
	}

	p.Saludar()
}
