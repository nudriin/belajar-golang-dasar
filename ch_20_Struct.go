package main

import "fmt"

type Person struct {
	// ! Penulisan propertiesnya menggunakan Pascal case
	Name    string
	Age     int
	City    string
	Address string
}

func main() {
	var person Person // menggunakann struct

	person.Name = "Nurdin"
	person.Age = 20
	person.City = "Sukamara"
	person.Address = "JL. Pangeran Samuder"

	fmt.Println(person)
	fmt.Println(person.Name)

	// ! CARA 2 mengunakan struct literals
	nurdin := Person{
		Name:    "Nurdin",
		Age:     20,
		City:    "Sukamara",
		Address: "JL. Pangeran Samudera",
	}

	fmt.Println(nurdin)

	// ! CARA 3 tetapi harus sesuai urutannya dengan properties yang ada pada struct nya
	sasuke := Person{"Sasuke", 20, "Sukamara", "JL. Sukarma"}
	fmt.Println(sasuke)
}
