package main

import "fmt"

type HasName interface {
	GetName() string
}

func interfaceImplements(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// ! IMPLEMENTASI 1
// *[1] Buat struct nya dulu ...1
type Persons struct {
	Name string
}

// ! IMPLEMENTASI 2
type Animal struct {
	Name string
}

// *[1] Kemudianbuat struct method yang sesuai dengan kontrak interfacenya ...2
func (person Persons) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Persons{
		Name: "Nurdin",
	}
	animal := Animal{
		Name: "Cat",
	}

	// *[1] tinggal inputkan struct objectnya ke function...3
	interfaceImplements(person)
	interfaceImplements(animal) // Bisa juga implementasi menggunakan struct animal
}
