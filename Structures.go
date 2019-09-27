package main

import (
	"fmt"
)

// Person : Structure for Person
type Person struct {
	Name string
	Age  int
}

// Male : Composition of Person
type Male struct {
	Person
	FacialHair bool
}

// Doctor : Structure for Doctor
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

// Structures : Basics of Structures
func Structures() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Aswin",
		Companions: []string{
			"Adam",
			"Joan",
			"Virat",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.ActorName)

	//Anonymous Struct
	bDoctor := struct {
		Name string
		Age  int
	}{Name: "John", Age: 12}
	fmt.Println(bDoctor)

	male := Male{
		Person:     Person{Name: "aswin", Age: 20},
		FacialHair: false,
	}
	male.Name = "Aswin VB"
	male.Age = 21
	fmt.Println(male)
}
