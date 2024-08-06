package main

import "fmt"

func main() {
	// ! Membuat constant
	const name = "Nurdin Hishasy"
	const age int = 20
	fmt.Println(name)
	fmt.Println(age)

	// ! Multiple constant
	const (
		VERSION = "1.1.0"
		AUTHOR  = "Nurdin"
	)

	fmt.Println(VERSION)
	fmt.Println(AUTHOR)

}
