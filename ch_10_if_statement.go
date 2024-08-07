package main

import "fmt"

func main() {
	ifElse()
	ifElseIf()
	ifElseWithShortStatement()
}

func ifElse() {
	name := "Nurdin"

	fmt.Println()
	fmt.Println("IF ELSE")
	// ! Bisa begini
	if name == "Nurdin" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hello Guest")
	}

	name = "Naruto"
	if name == "Nurdin" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hello Guest")
	}
}

func ifElseIf() {
	name := "Naruto"
	fmt.Println()
	fmt.Println("IF ELSE IF")
	if name == "Nurdin" {
		fmt.Println("Hello Nurdin")
	} else if name == "Naruto" {
		fmt.Println("Hello Naruto")
	} else {
		fmt.Println("Hello Guest")
	}
}

func ifElseWithShortStatement() {
	name := "Nurdin"

	fmt.Println()
	fmt.Println("IF WITH SHORT STATEMENT")
	// * Kita dapat membuat short statement sebelum kondisinya
	if length := len(name); length > 6 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama benar")
	}
}
