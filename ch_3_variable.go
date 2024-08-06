package main

import (
	"fmt"
)

func main() {
	var name string // ! Declare variabel with var keyword

	name = "Nurdin"
	fmt.Println(name)

	name = "Nurdin Hishasy"
	fmt.Println(name)

	// !langsung di inisiasi dan akan di detect otomatis
	var city = "Sukamara"
	fmt.Println(city)

	// !langsung di inisiasi dan akan di detect otomatis
	age := 20
	fmt.Println(age)

	age = 30
	fmt.Println(age)

	// ! Multiple varibale (membuat banyak variable sekaligus)
	var (
		country    string
		address    string
		postalCode int
	)

	country = "Indonesia"
	address = "JL. Setia Yakin"
	postalCode = 74712

	fmt.Println(country)
	fmt.Println(address)
	fmt.Println(postalCode)
}
